package delivery

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/sdf0106/ip-project/internal/service"
	"net/http"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X_CSRF_Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Group(func(mux chi.Router) {
		mux.Route("/auth", func(mux chi.Router) {
			mux.Post("/sign_in", h.signIn)
			mux.Post("/sign_up", h.signUp)
		})
	})

	mux.Group(func(mux chi.Router) {
		mux.Use(h.authMiddleware)

		mux.Route("/user", func(mux chi.Router) {
			mux.Post("/choose_role", h.chooseRole)
			mux.Put("/change_role", h.updateRole)
			mux.Put("/update_user_info", h.updateUserInfo)
		})

		mux.Route("/agent", func(mux chi.Router) {
			mux.Get("/", h.getAllAgents)
			mux.Get("/{id}", h.getAgentById)
			mux.Get("/houses", h.getAllHouses)
		})

		mux.Route("/house", func(mux chi.Router) {
			mux.Get("/", h.getAllHouses)
			mux.Get("/{id}", h.getHouseById)
		})

		mux.Route("/owner", func(mux chi.Router) {
			mux.Get("/house", h.getMyHouses)
			mux.Post("/house", h.createHouse)
			mux.Delete("/house/{id}", h.deleteHouse)
			mux.Put("/house", h.updateHouse)
			mux.Post("/agent", h.hireAgent)
		})

		mux.Route("/client", func(mux chi.Router) {
			//TODO: Add additional endpoints to get client info
			mux.Route("/cart", func(mux chi.Router) {
				mux.Get("/cart", h.getCart)

				mux.Route("/cart/item", func(mux chi.Router) {
					mux.Post("/", h.addToCart)
					mux.Delete("/{id}", h.removeFromCart)
				})
			})

		})
	})

	return mux
}
