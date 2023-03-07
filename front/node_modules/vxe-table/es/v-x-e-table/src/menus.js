import VXEStore from './store';
export var menus = new VXEStore();
if (process.env.NODE_ENV === 'development') {
    Object.assign(menus, { _name: 'Menus' });
}
