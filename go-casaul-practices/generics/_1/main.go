package main

import (
	"fmt"
)

type DTOType int

const (
	User DTOType = iota
	Product
)

// Then use these constants in Register and Create calls

// BaseDTO 是所有 DTO 的基礎接口
type BaseDTO interface {
	print() string
}

// Creator 是創建 DTO 的函數類型
type Creator func() BaseDTO

// DTOFactory 管理 DTO 的創建
type DTOFactory struct {
	// creators map[string]Creator
	creators map[DTOType]Creator
}

// NewDTOFactory 創建一個新的 DTOFactory
func NewDTOFactory() *DTOFactory {
	return &DTOFactory{
		// creators: make(map[string]Creator),
		creators: make(map[DTOType]Creator),
	}
}

func (f *DTOFactory) Register(dtoType DTOType, creator Creator) {
	f.creators[dtoType] = creator
}

func (f *DTOFactory) Create(dtoType DTOType) (BaseDTO, error) {
	creator, exists := f.creators[dtoType]
	if !exists {
		return nil, fmt.Errorf("unsupported DTO type: %s", dtoType)
	}
	return creator(), nil
}

// UserDTO 示例
type UserDTO struct{}

func (u UserDTO) print() string { return "user" }

// ProductDTO 示例
type ProductDTO struct{}

func (p ProductDTO) print() string { return "product" }

func main() {
	factory := NewDTOFactory()

	// 註冊創建函數
	factory.Register(User, func() BaseDTO { return &UserDTO{} })
	factory.Register(Product, func() BaseDTO { return &ProductDTO{} })

	// 創建 DTO
	userDTO, _ := factory.Create(User)
	productDTO, _ := factory.Create(Product)

	fmt.Printf("Created UserDTO: %T\n", userDTO)
	fmt.Printf("Created ProductDTO: %T\n", productDTO)
}
