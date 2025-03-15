package product

type Product struct {
	ID          uint
	Nama        string
	Brand       string
	Category    string
	Price       uint
	Amount      uint
	Description string
	Image       string
}

type ProductService interface {
	CreateProduct(newData Product) (Product, error)
	GetAllProducts(offset int, category, brand, sort, q string) ([]Product, error)
	GetProductByID(productID uint) (Product, error)
	// UpdateProductByID(productid uint, newData Product) (Product, error)
	// DeleteProductByID(productid uint) error
}

type ProductModel interface {
	CreateProduct(newData Product) (Product, error)
	GetAllProducts(offset int, category, brand, sort, q string) ([]Product, error)
	GetProductByID(productID uint) (Product, error)
	// UpdateProductByID(productid uint, newData Product) (Product, error)
	// DeleteProductByID(productid uint) error
}