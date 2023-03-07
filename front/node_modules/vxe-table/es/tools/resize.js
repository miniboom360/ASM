import XEUtils from 'xe-utils';
import GlobalConfig from '../v-x-e-table/src/conf';
/**
 * 监听 resize 事件
 * 如果项目中已使用了 resize-observer-polyfill，那么只需要将方法定义全局，该组件就会自动使用
 */
var resizeTimeout;
/* eslint-disable no-use-before-define */
var eventStore = [];
var defaultInterval = 500;
function eventHandle() {
    if (eventStore.length) {
        eventStore.forEach(function (item) {
            item.tarList.forEach(function (observer) {
                var target = observer.target, width = observer.width, heighe = observer.heighe;
                var clientWidth = target.clientWidth;
                var clientHeight = target.clientHeight;
                var rWidth = clientWidth && width !== clientWidth;
                var rHeight = clientHeight && heighe !== clientHeight;
                if (rWidth || rHeight) {
                    observer.width = clientWidth;
                    observer.heighe = clientHeight;
                    setTimeout(item.callback);
                }
            });
        });
        /* eslint-disable @typescript-eslint/no-use-before-define */
        eventListener();
    }
}
function eventListener() {
    clearTimeout(resizeTimeout);
    resizeTimeout = setTimeout(eventHandle, GlobalConfig.resizeInterval || defaultInterval);
}
var XEResizeObserver = /** @class */ (function () {
    function XEResizeObserver(callback) {
        this.tarList = [];
        this.callback = callback;
    }
    XEResizeObserver.prototype.observe = function (target) {
        var _this = this;
        if (target) {
            var tarList = this.tarList;
            if (!tarList.some(function (observer) { return observer.target === target; })) {
                tarList.push({
                    target: target,
                    width: target.clientWidth,
                    heighe: target.clientHeight
                });
            }
            if (!eventStore.length) {
                eventListener();
            }
            if (!eventStore.some(function (item) { return item === _this; })) {
                eventStore.push(this);
            }
        }
    };
    XEResizeObserver.prototype.unobserve = function (target) {
        XEUtils.remove(eventStore, function (item) { return item.tarList.some(function (observer) { return observer.target === target; }); });
    };
    XEResizeObserver.prototype.disconnect = function () {
        var _this = this;
        XEUtils.remove(eventStore, function (item) { return item === _this; });
    };
    return XEResizeObserver;
}());
export { XEResizeObserver };
export function createResizeEvent(callback) {
    if (window.ResizeObserver) {
        return new window.ResizeObserver(callback);
    }
    return new XEResizeObserver(callback);
}
