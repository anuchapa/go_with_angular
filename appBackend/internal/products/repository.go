package products

import (
	"context"
	"goBackend/internal/adapter/db/model"
	"goBackend/internal/adapter/db/query"
	"log"

	"gorm.io/gorm"
)

type Repository interface {
	FindByID(ctx context.Context, id int64) (*model.Product, error)
	FindAll(ctx context.Context) ([]*model.Product, error)
	Create(ctx context.Context, products []*model.Product) error
	Delete(ctx context.Context, id int64) (int64,error)
}

type repository struct {
	query *query.Query
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{query: query.Use(db)}
}

func (r *repository) FindByID(ctx context.Context, id int64) (*model.Product, error) {
	product, err := r.query.Product.WithContext(ctx).Where(query.Product.ID.Eq(id)).First()
	if err != nil {
		log.Printf("failed to query product by id: %d, err:%s.", id, err.Error())
		return nil, err
	}
	return product, nil
}
func (r *repository) FindAll(ctx context.Context) ([]*model.Product, error) {
	products, err := r.query.Product.WithContext(ctx).Find()

	if err != nil {
		log.Printf("failed to query products, err:%s.", err.Error())
		return nil, err
	}
	return products, nil
}
func (r *repository) Create(ctx context.Context, products []*model.Product) error {
	tx := r.query.Begin()
	defer tx.Rollback()

	if err := tx.Product.WithContext(ctx).Create(products...); err != nil {
		log.Printf("Created is failed, err:%s", err.Error())
		return err
	}

	if err := tx.Commit(); err != nil {
		log.Printf("Commit products created is failed, err : %s", err.Error())
		return err
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, id int64) (int64,error) {
	tx := r.query.Begin()
	info, err := tx.Product.WithContext(ctx).Delete(&model.Product{ID: id})
	defer tx.Rollback()
	if err != nil {
		log.Printf("failed to query product by id: %d, err:%s.", id, err.Error())
		return 0,err
	}

	if info.RowsAffected == 0 {
		log.Printf("RowsAffected:0 , id: %d not fount",id)
		return info.RowsAffected,nil
	}

	tx.Commit()
	return info.RowsAffected,nil
}
