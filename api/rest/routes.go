package main

import (
  "github.com/gorilla/mux"
  "tracksale.prova/estruturas/route"
)

type Routes []route.Route

func NewRouter() *mux.Router {

  router := mux.NewRouter().StrictSlash(true)
  for _, route := range routes {
      router.
          Methods(route.Method).
          Path(route.Pattern).
          Name(route.Name).
          Handler(route.HandlerFunc)
  }

  return router
}

var routes = Routes{
  route.Route{
      "Index",
      "GET",
      "/",
      Index,
  },
  route.Route{
      "TodoIndex",
      "GET",
      "/todos",
      TodoIndex,
  },
  route.Route{
      "TodoShow",
      "GET",
      "/todos/{todoId}",
      TodoShow,
  },
  route.Route{
    "Status",
    "POST",
    "/status",
    MostraStatusPi,
  },
}
