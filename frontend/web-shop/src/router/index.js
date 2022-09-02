import { createWebHistory, createRouter } from "vue-router";
import Login from "../components/Login.vue"
import Counter from "../components/Counter.vue"
import MovieList from "../components/MovieList.vue"


const routes = [
    {
        path: "/login",
        name: Login,
        component: Login
    },
    {
        path: "/counter",
        name: Counter,
        component: Counter
    },
    {
        path: "/movies",
        name: MovieList,
        component: MovieList
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes: routes,
})

export default router;