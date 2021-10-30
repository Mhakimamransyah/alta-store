package productsimages

import (
	productsimages "altaStore/business/products_images"
	"altaStore/modules/products"
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	awsSession "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"math/rand"
	"mime/multipart"
	"os"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type GormRepository struct {
	DB *gorm.DB
}

const (
	folder    = "products_image"
	S3_REGION = "ap-southeast-1"
	S3_BUCKET = "alta-store-bucket"
)

type ProductsImagesTable struct {
	ID          int                    `gorm:"id;primaryKey:autoIncrement"`
	FileName    string                 `gorm:"image_path"`
	Products_ID int                    `gorm:"id_products"`
	Created_at  time.Time              `gorm:"created_at;type:datetime;default:null"`
	Updated_at  time.Time              `gorm:"updated_at;type:datetime;default:null"`
	Deleted_at  time.Time              `gorm:"deleted_at;type:datetime;default:null"`
	Products    products.ProductsTable `gorm:"foreignKey:Products_ID"`
}

func ConvertProductsImagesToProductsImagesTable(product_images *productsimages.ProductImages) *ProductsImagesTable {
	return &ProductsImagesTable{
		ID:          product_images.ID,
		FileName:    product_images.FileName,
		Products_ID: product_images.Products_ID,
		Created_at:  product_images.Created_at,
		Updated_at:  product_images.Updated_at,
		Deleted_at:  product_images.Deleted_at,
	}
}

func ConvertProductsImagesTableToProductImages(productsimagestables *ProductsImagesTable) *productsimages.ProductImages {
	return &productsimages.ProductImages{
		ID:          productsimagestables.ID,
		FileName:    productsimagestables.FileName,
		Products_ID: productsimagestables.Products_ID,
		Created_at:  productsimagestables.Created_at,
		Updated_at:  productsimagestables.Updated_at,
		Deleted_at:  productsimagestables.Deleted_at,
		Path:        os.Getenv("ALTASTORE_S3_ADDRESS") + productsimagestables.FileName,
	}
}

func InitProductsImagesRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		DB: db,
	}
}

func (repository *GormRepository) CreateImages(products_images *productsimages.ProductImages, files *multipart.FileHeader, createdBy int) error {
	filename := products_images.FileName + "-" + strconv.Itoa(time.Now().Second()) + "-" + strconv.Itoa(time.Now().Minute()) + strconv.Itoa(rand.Intn(1000)) + ".png"
	products_images_table := ConvertProductsImagesToProductsImagesTable(products_images)
	products_images_table.FileName = filename
	err := repository.DB.Save(products_images_table).Error
	if err != nil {
		return err
	}
	UploadToS3(files, filename)

	return nil
}

func UploadToS3(files *multipart.FileHeader, filename string) error {
	sess, err := awsSession.NewSession(&aws.Config{
		Region: aws.String(S3_REGION),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("ALTASTORE_S3_ID"),
			os.Getenv("ALTASTORE_S3_KEY"),
			"",
		),
	})

	if err != nil {
		return err
	}

	src, err := files.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	buffer := make([]byte, files.Size)

	src.Read(buffer)

	_, err = s3.New(sess).PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(S3_BUCKET),
		Key:         aws.String("products-images/" + filename),
		Body:        bytes.NewReader(buffer),
		ACL:         aws.String("public-read"),
		ContentType: aws.String("image/jpeg"),
	})

	if err != nil {
		return err
	}
	return nil
}

func (repository *GormRepository) GetListProductsImagesByIdProducts(id_products int) (*[]productsimages.ProductImages, error) {
	var list_products_images_tables []ProductsImagesTable
	err := repository.DB.Where("products_id = ?", id_products).Find(&list_products_images_tables).Error
	if err != nil {
		return nil, err
	}

	var list_products_images []productsimages.ProductImages
	for _, data := range list_products_images_tables {
		list_products_images = append(list_products_images, *ConvertProductsImagesTableToProductImages(&data))
	}

	return &list_products_images, nil
}

func (repository *GormRepository) DeleteProductsImages(products_images *productsimages.ProductImages, deletedById int) error {
	products_images_table := ConvertProductsImagesToProductsImagesTable(products_images)
	path := folder + "/" + products_images_table.FileName
	err := repository.DB.Where("id = ?", products_images.ID).Delete(&products_images_table).Error
	if err != nil {
		return err
	}

	// remove file system
	err = os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}

func (repository *GormRepository) GetProductsImagesById(id_products_image int) (*productsimages.ProductImages, error) {
	var products_images_table ProductsImagesTable
	err := repository.DB.Where("id = ?", id_products_image).First(&products_images_table).Error
	if err != nil {
		return nil, err
	}
	return ConvertProductsImagesTableToProductImages(&products_images_table), nil
}
