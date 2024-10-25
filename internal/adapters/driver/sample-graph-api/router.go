package samplegraphapi

import (
	"context"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/nphiro/hexa/internal/adapters/driver/sample-graph-api/gen"
	"github.com/vektah/gqlparser/v2/ast"
)

func New() http.Handler {
	resolver := &Resolver{}
	execSchema := gen.NewExecutableSchema(gen.Config{Resolvers: resolver})

	h := handler.New(execSchema)
	{
		h.AddTransport(transport.SSE{})
		h.AddTransport(transport.Options{})
		h.AddTransport(transport.GET{})
		h.AddTransport(transport.POST{})
		h.AddTransport(transport.MultipartForm{})

		h.SetQueryCache(lru.New[*ast.QueryDocument](1000))
		h.Use(extension.Introspection{})
		h.Use(extension.AutomaticPersistedQuery{
			Cache: lru.New[string](100),
		})

		h.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
			graphql.GetOperationContext(ctx).DisableIntrospection = false
			return next(ctx)
		})
	}

	r := gin.New()
	r.POST("/query", gin.WrapH(h))
	r.GET("/playground", gin.WrapH(playground.Handler("GraphQL playground", "/query")))

	return r
}
