const urls = {
  authBaseUrl: '/auth',
  apiBaseUrl: '/api'
};

const token = {
  prefix: 'Bearer ',
  cookieName: 'AccessToken',
  refreshFlag: 'TokenRefreshFlag',
  expiredTime: '1h',
  maxRefreshTime: '48h'
};

const form = {
  labelCol: {span: 3},
  wrapperCol: {span: 15},
  submitCol: {span: 15, push: 3}
};

const fluidForm = {
  labelCol: {span: 3},
  wrapperCol: {span: 21},
  submitCol: {span: 21, push: 3}
};

const navs = [
  {label: '美图', icon: 'icon-meitu', to: '/picture', key: 'picture'},
  {label: '视频', icon: 'icon-shipin', to: '/video', key: 'video'},
  {label: '电影', icon: 'icon-dy-light-copy', to: '/movie', key: 'movie'},
];

const userMenu = [
  {label: '个人中心', icon: 'idcard', to: '/profile'},
  {label: '系统设置', icon: 'setting', to: '/settings'},
  {label: '注销登录', icon: 'logout', to: '/logout'},
];

export default {
  urls, token, form, fluidForm, navs, userMenu
}