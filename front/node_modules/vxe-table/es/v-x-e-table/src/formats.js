import VXEStore from './store';
export var formats = new VXEStore();
if (process.env.NODE_ENV === 'development') {
    Object.assign(formats, { _name: 'Formats' });
}
