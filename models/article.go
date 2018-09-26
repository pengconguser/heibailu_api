package models

type Article struct {
	BaseModel

	Title       string `json:"title"`
	Author      string `json:"author"`
	UserId      int    `json:"user_id"`
	Hits        int    `json:"hits"`
	Content     string `json:"content"`
	CategoryId  int    `json:"category_id"`
	Description string `json:"description"`
}

func (self *Article) _GetId(id string) error {
	var article Article

	err := DB.First(&article, id).Error

	if err != nil {
		return err
	}

	return nil
}

func GetIdArticle(id string) (Article, error) {
	var article Article

	err := article._GetId(id)

	if err != nil {
		return article, err
	}

	return article, nil
}
