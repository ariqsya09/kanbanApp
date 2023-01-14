package api

import (
	"a21hc3NpZ25tZW50/entity"
	"a21hc3NpZ25tZW50/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type CategoryAPI interface {
	GetCategory(w http.ResponseWriter, r *http.Request)
	CreateNewCategory(w http.ResponseWriter, r *http.Request)
	DeleteCategory(w http.ResponseWriter, r *http.Request)
	GetCategoryWithTasks(w http.ResponseWriter, r *http.Request)
}

type categoryAPI struct {
	categoryService service.CategoryService
}

func NewCategoryAPI(categoryService service.CategoryService) *categoryAPI {
	return &categoryAPI{categoryService}
}

func (c *categoryAPI) GetCategory(w http.ResponseWriter, r *http.Request) {
	// var category []entity.Category

	// id := r.Context().Value("id")

	// if id.(string) == "" {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
	// 	return
	// }

	// cat, err := c.categoryService.GetCategories(r.Context(), id.(int))
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
	// 	return
	// }

	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(cat)
	// return

	userId := r.Context().Value("id")

	idLogin, err := strconv.Atoi(userId.(string))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("get category task", err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}

	categories, err := c.categoryService.GetCategories(r.Context(), idLogin)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)

	// TODO: answer here
}

func (c *categoryAPI) CreateNewCategory(w http.ResponseWriter, r *http.Request) {
	var category entity.CategoryRequest

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid category request"))
		return
	}

	if category.Type == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid category request"))
		return
	}

	userId := r.Context().Value("id")
	idLogin, err := strconv.Atoi(userId.(string))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}
	dataCat := entity.Category{
		Type:   category.Type,
		UserID: idLogin,
	}
	saveCat, err := c.categoryService.StoreCategory(r.Context(), &dataCat)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}

	var resp = map[string]interface{}{}
	resp["user_id"] = saveCat.UserID
	resp["category_id"] = saveCat.ID
	resp["message"] = "success create new category"
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&resp)

	return
	// TODO: answer here
}

func (c *categoryAPI) DeleteCategory(w http.ResponseWriter, r *http.Request) {

	userId := r.Context().Value("id")
	idLogin, err := strconv.Atoi(userId.(string))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("get category task", err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}

	categoryID := r.URL.Query().Get("category_id")

	categoryIDConv, _ := strconv.Atoi(categoryID)

	err = c.categoryService.DeleteCategory(r.Context(), categoryIDConv)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}

	var resp = map[string]interface{}{}
	resp["user_id"] = idLogin
	resp["category_id"] = categoryID
	resp["message"] = "success delete category"
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
	return
	// TODO: answer here
}

func (c *categoryAPI) GetCategoryWithTasks(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("id")
	idLogin, err := strconv.Atoi(userId.(string))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("get category task", err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}

	categories, err := c.categoryService.GetCategoriesWithTasks(r.Context(), int(idLogin))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("internal server error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)

}
