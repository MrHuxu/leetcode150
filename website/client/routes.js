import Index from './components/index';
import Detail from './components/detail';
import NoMatch from './components/404';

export default [
  { exact: true, path: '/', component: Index },
  { path: '/:id', component: Detail },

  { path: '/404', component: NoMatch }
];
