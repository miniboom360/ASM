import { watch } from 'vue';
import XEUtils from 'xe-utils';
import { OptionInfo } from './optionInfo';
export function isOption(option) {
    return option instanceof OptionInfo;
}
export function createOption($xeselect, _vm) {
    return isOption(_vm) ? _vm : new OptionInfo($xeselect, _vm);
}
export function watchOption(props, option) {
    Object.keys(props).forEach(function (name) {
        watch(function () { return props[name]; }, function (value) {
            option.update(name, value);
        });
    });
}
export function assemOption($xeselect, el, option, optgroup) {
    var reactData = $xeselect.reactData;
    var staticOptions = reactData.staticOptions;
    var parentElem = el.parentNode;
    var parentOption = optgroup ? optgroup.option : null;
    var parentCols = parentOption ? parentOption.options : staticOptions;
    if (parentElem && parentCols) {
        parentCols.splice(XEUtils.arrayIndexOf(parentElem.children, el), 0, option);
        reactData.staticOptions = staticOptions.slice(0);
    }
}
export function destroyOption($xeselect, option) {
    var reactData = $xeselect.reactData;
    var staticOptions = reactData.staticOptions;
    var matchObj = XEUtils.findTree(staticOptions, function (item) { return item.id === option.id; }, { children: 'options' });
    if (matchObj) {
        matchObj.items.splice(matchObj.index, 1);
    }
    reactData.staticOptions = staticOptions.slice(0);
}
