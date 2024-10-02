import { createRouter, createWebHistory } from "vue-router";
import Home from "./views/Home.vue";
import About from "./views/About.vue";
import Contact from "./views/Contact.vue";

const routes = [
    { path: "/", component: Home, name: "Home" },
    { path: "/about", component: About, name: "About" },
    { path: "/contact", component: Contact, name: "Contact" },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;
