import VXEStore from './store';
export var commands = new VXEStore();
if (process.env.NODE_ENV === 'development') {
    Object.assign(commands, { _name: 'Commands' });
}
