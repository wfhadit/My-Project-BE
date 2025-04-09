package services

import (
	"context"
	"errors"
	"log"
	"mime/multipart"
	"my-project-be/config"
	cart "my-project-be/features/cart"
	"my-project-be/features/order"
	user "my-project-be/features/user"
	"my-project-be/features/user/handler"
	"my-project-be/helper"
	"my-project-be/lib/cloudinary"
	"my-project-be/middlewares"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/veritrans/go-midtrans"
)

type service struct {
	cart  cart.CartModel
	model user.UserModel
	pm helper.PasswordManager
	v *validator.Validate
	modelOrder order.OrderModel
	midtrans midtrans.Client
}

func NewService(m user.UserModel, c cart.CartModel, mo order.OrderModel,midtrans midtrans.Client) user.UserService {
	return &service{
		model: m,
		pm: helper.NewPasswordManager(),
		v: validator.New(),
		cart: c,
		modelOrder: mo,
		midtrans: midtrans,
	}
}

func (s *service) Register(newData user.User) error {
	
	registerValidate := handler.RegisterRequest{
		Nama:     newData.Nama,
		Email:    newData.Email,
		Password: newData.Password,
	}
	err := s.v.Struct(&registerValidate)
	if err != nil {
		return err
	}
	exist, errCheck := s.model.Login(newData.Email)
	if errCheck == nil && exist.ID != 0 {
		return errors.New("Email already exist")
	}

	result, _ := s.pm.HashPassword(newData.Password)
	newData.Password = result
	newData.Foto = "https://res.cloudinary.com/dvehysudh/image/upload/kmls5vfsijivozf8elib.jpg"

	err = s.model.Register(newData)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) Login(loginData user.User) (user.User, string, []cart.Cart,[]order.Order,  error) {
	loginValidate :=handler.LoginRequest{
		Email:    loginData.Email,
		Password: loginData.Password,
	}
	err := s.v.Struct(&loginValidate)
	if err != nil {
		log.Println("terjadi error", err.Error())
		return user.User{}, "",[]cart.Cart{},[]order.Order{}, err
	}

	data, err := s.model.Login(loginValidate.Email)
	if err != nil {
		return user.User{}, "", []cart.Cart{},[]order.Order{},err
	}

	err = s.pm.CheckPassword(loginValidate.Password, data.Password)
	if err != nil {
		return user.User{}, "", []cart.Cart{},[]order.Order{},err
	}

	token, err := middlewares.GenerateJWT(data.ID, data.Nama)
	if err != nil {
		return user.User{}, "", []cart.Cart{},[]order.Order{},err
	}
	cartData, err := s.cart.GetCart(data.ID)
	if err != nil {
		log.Println("error dari database cart", err.Error())
		return user.User{},"", []cart.Cart{}, []order.Order{},err
	}
	orders, err := s.modelOrder.GetAllOrders(data.ID)
	if err != nil {
		log.Println("error dari database order", err.Error())
		return user.User{},"", []cart.Cart{}, []order.Order{},err
	}
	client := s.midtrans
	coreGateway := midtrans.CoreGateway{Client: client}
	updatedOrders := []order.Order{}
	for _, order := range orders {
		if order.Status == "pending" {
			resp, err := coreGateway.Status(order.OrderUniqueID)
			if err != nil {
				log.Println(err.Error())
				continue
			}
			if order.Status != resp.TransactionStatus {order.Status = resp.TransactionStatus
				_, errUpdate := s.modelOrder.GetOrderByUniqueID(order.OrderUniqueID, data.ID, order.Status)
				if errUpdate != nil {
					log.Println(errUpdate.Error())
				}
			}
		}
		updatedOrders = append(updatedOrders, order)
	}


	return data, token, cartData,updatedOrders,nil
}

func (s *service) KeepLogin(token *jwt.Token) (user.User, string, []cart.Cart,[]order.Order,error) {
	userID,_:= middlewares.DecodeToken(token)
	result, err := s.model.GetUserByID(userID)
	if err != nil {
		log.Println("error dari database user", err.Error())
		return user.User{},  "",[]cart.Cart{},[]order.Order{},err
	}
	newToken, err := middlewares.GenerateJWT(result.ID, result.Nama)
	if err != nil {
		log.Println("error dari token baru", err.Error())
		return user.User{},"", []cart.Cart{},[]order.Order{}, err
	}
	cartData, err := s.cart.GetCart(userID)
	if err != nil {
		log.Println("error dari database cart", err.Error())
		return user.User{},"", []cart.Cart{},[]order.Order{}, err
	}
	orders, err := s.modelOrder.GetAllOrders(userID)
	if err != nil {
		log.Println("error dari database order", err.Error())
		return user.User{},"", []cart.Cart{},[]order.Order{}, err
	}
	client := s.midtrans
	coreGateway := midtrans.CoreGateway{Client: client}
	updatedOrders := []order.Order{}
	for _, order := range orders {
		if order.Status == "pending" {
			resp, err := coreGateway.Status(order.OrderUniqueID)
			if err != nil {
				log.Println(err.Error())
				continue
			}
			if order.Status != resp.TransactionStatus {order.Status = resp.TransactionStatus
				_, errUpdate := s.modelOrder.GetOrderByUniqueID(order.OrderUniqueID, userID, order.Status)
				if errUpdate != nil {
					log.Println(errUpdate.Error())
				}
			}
		}
		updatedOrders = append(updatedOrders, order)
	}


	return result, newToken,cartData, updatedOrders,nil
}

func (s *service) Update(token *jwt.Token, newData user.User, file *multipart.FileHeader) (user.User, error) {
	userID,_ := middlewares.DecodeToken(token)
	existingUser, err := s.model.GetUserByID(userID)
	if err != nil {
		return user.User{}, errors.New("user tidak ditemukan")
	}
	if newData.Nama != "" {
		existingUser.Nama = newData.Nama
	}
	if newData.Email != "" {
		existingUser.Email = newData.Email
	}
	if newData.TanggalLahir != "" && newData.TanggalLahir != "undefined" {
		existingUser.TanggalLahir = newData.TanggalLahir
	}
	if newData.JenisKelamin != "" {
		existingUser.JenisKelamin = newData.JenisKelamin
	}
	if newData.NomorHP != "" {
		existingUser.NomorHP = newData.NomorHP
	}
	if newData.Alamat != "" {
		existingUser.Alamat = newData.Alamat
	}
	if file != nil {

		Buka, err := file.Open()
	 	if err != nil {
		return user.User{}, errors.New("file tidak terbaca")
		}
		defer Buka.Close()

		cfg := config.InitConfig()
		cld, err := cloudinary.GetCloudinaryClient(&cfg)
		if err != nil {
		return user.User{}, errors.New("gagal autentikasi cld")
		}

		uploadResult, err := cld.Upload.Upload(context.Background(), Buka, uploader.UploadParams{})
		if err != nil {
		return user.User{}, errors.New("gagal upload ke cloudinary")
		}
	
		existingUser.Foto = uploadResult.SecureURL
	}

	result, err := s.model.Update(userID, existingUser)
	if err != nil {
		return user.User{}, err
	}

	return result, nil
}