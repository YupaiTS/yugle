import Vue from 'vue'
import Router from 'vue-router'
import 'vue-cookies'
import consts from '../consts'

import Index from './views/Index'

Vue.use(Router)

const router = new Router({
  routes: [
    {
      path: '/',
      component: Index
    },
  ]
})

router.beforeEach((to, from, next) => {
  const accessToken = window.$cookies.get(consts.token.cookieName);
  const tokenValid = accessToken !== undefined && accessToken !== null && accessToken.trim() !== '';
  if (tokenValid) {
    next();
  } else {
    window.location.href = '/login';
  }
})

export default router