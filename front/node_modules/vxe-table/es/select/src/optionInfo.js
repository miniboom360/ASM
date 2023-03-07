import XEUtils from 'xe-utils';
var OptionInfo = /** @class */ (function () {
    function OptionInfo($xeselect, _vm) {
        Object.assign(this, {
            id: XEUtils.uniqueId('option_'),
            value: _vm.value,
            label: _vm.label,
            visible: _vm.visible,
            className: _vm.className,
            disabled: _vm.disabled
        });
    }
    OptionInfo.prototype.update = function (name, value) {
        this[name] = value;
    };
    return OptionInfo;
}());
export { OptionInfo };
