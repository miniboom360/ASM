import XEUtils from 'xe-utils';
import { browse } from './dom';
export var EVENT_KEYS = {
    F2: 'F2',
    ESCAPE: 'Escape',
    ENTER: 'Enter',
    TAB: 'Tab',
    DELETE: 'Delete',
    BACKSPACE: 'Backspace',
    SPACEBAR: ' ',
    CONTEXT_MENU: 'ContextMenu',
    ARROW_UP: 'ArrowUp',
    ARROW_DOWN: 'ArrowDown',
    ARROW_LEFT: 'ArrowLeft',
    ARROW_RIGHT: 'ArrowRight',
    PAGE_UP: 'PageUp',
    PAGE_DOWN: 'PageDown'
};
var convertEventKeys = {
    ' ': 'Spacebar',
    Apps: EVENT_KEYS.CONTEXT_MENU,
    Del: EVENT_KEYS.DELETE,
    Up: EVENT_KEYS.ARROW_UP,
    Down: EVENT_KEYS.ARROW_DOWN,
    Left: EVENT_KEYS.ARROW_LEFT,
    Right: EVENT_KEYS.ARROW_RIGHT
};
// 监听全局事件
var wheelName = browse.firefox ? 'DOMMouseScroll' : 'mousewheel';
var eventStore = [];
export var hasEventKey = function (evnt, targetKey) {
    var key = evnt.key;
    targetKey = targetKey.toLowerCase();
    return key ? (targetKey === key.toLowerCase() || !!(convertEventKeys[key] && convertEventKeys[key].toLowerCase() === targetKey)) : false;
};
function triggerEvent(evnt) {
    var isWheel = evnt.type === wheelName;
    eventStore.forEach(function (_a) {
        var type = _a.type, cb = _a.cb;
        // 如果被取消冒泡，不再执行
        if (!evnt.cancelBubble) {
            if (type === evnt.type || (isWheel && type === 'mousewheel')) {
                cb(evnt);
            }
        }
    });
}
export var GlobalEvent = {
    on: function (comp, type, cb) {
        eventStore.push({ comp: comp, type: type, cb: cb });
    },
    off: function (comp, type) {
        XEUtils.remove(eventStore, function (item) { return item.comp === comp && item.type === type; });
    },
    trigger: triggerEvent,
    eqKeypad: function (evnt, keyVal) {
        var key = evnt.key;
        if (keyVal.toLowerCase() === key.toLowerCase()) {
            return true;
        }
        return false;
    }
};
if (browse.isDoc) {
    if (!browse.msie) {
        document.addEventListener('copy', triggerEvent, false);
        document.addEventListener('cut', triggerEvent, false);
        document.addEventListener('paste', triggerEvent, false);
    }
    document.addEventListener('keydown', triggerEvent, false);
    document.addEventListener('contextmenu', triggerEvent, false);
    window.addEventListener('mousedown', triggerEvent, false);
    window.addEventListener('blur', triggerEvent, false);
    window.addEventListener('resize', triggerEvent, false);
    window.addEventListener(wheelName, XEUtils.throttle(triggerEvent, 100, { leading: true, trailing: false }), { passive: true, capture: false });
}
