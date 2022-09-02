

import { createStore } from "vuex";

const store = createStore({
  state: {
    movies: [
        {
            id: 1,
            title: "Avatar",
            description:"Jake Sully lives with his newfound family formed on the planet of Pandora. Once a familiar threat returns to finish what was previously started, Jake must work with Neytiri and the army of the Na'vi race to protect their planet.",
            price:7.0,
            poster: "https://cdn.shopify.com/s/files/1/0057/3728/3618/products/280186877_327463976185980_4928049883994991289_n_480x.progressive.jpg?v=1653311094"
        },
        {
            id: 1,
            title: "Thor: Love and Thunder",
            description:"Thor enlists the help of Valkyrie, Korg and ex-girlfriend Jane Foster to fight Gorr the God Butcher, who intends to make the gods extinct.",
            price:12.0,
            poster: "https://cdn.shopify.com/s/files/1/0057/3728/3618/products/thor_480x.progressive.jpg?v=1653571052"
        },
        {
            id: 1,
            title: "Batman",
            description:"When a sadistic serial killer begins murdering key political figures in Gotham, Batman is forced to investigate the city's hidden corruption and question his family's involvement.",
            price:4.0,
            poster: "https://cdn.shopify.com/s/files/1/0057/3728/3618/products/the-batman_tgstxmov_480x.progressive.jpg?v=1641930817"
        },
    ] ,
    cart: [],
    token: ''
  },
  mutations:{
    addCartItem(state, movie){
        state.cart.push(movie)
    },
    removeCartItem(state, movie) {
      state.cart = state.cart.filter((cartItem) => {
        return cartItem.id != movie.id;
      });
    },
    addApiToken(state, token){
        state.token = token
    },
    removeApiToken(state, token){
        state.token = token
    }
  },
  getters:{
    getToken(state){
        return state.token
    }
  }
})

export default store;