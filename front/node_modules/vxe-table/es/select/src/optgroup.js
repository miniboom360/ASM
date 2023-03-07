import { defineComponent, h, onUnmounted, provide, inject, ref, onMounted } from 'vue';
import { createOption, watchOption, destroyOption, assemOption } from './util';
export default defineComponent({
    name: 'VxeOptgroup',
    props: {
        label: { type: [String, Number, Boolean], default: '' },
        visible: { type: Boolean, default: null },
        className: [String, Function],
        disabled: Boolean
    },
    setup: function (props, _a) {
        var slots = _a.slots;
        var elem = ref();
        var $xeselect = inject('$xeselect', {});
        var option = createOption($xeselect, props);
        var xeoption = { option: option };
        option.options = [];
        provide('xeoptgroup', xeoption);
        watchOption(props, option);
        onMounted(function () {
            assemOption($xeselect, elem.value, option);
        });
        onUnmounted(function () {
            destroyOption($xeselect, option);
        });
        return function () {
            return h('div', {
                ref: elem
            }, slots.default ? slots.default() : []);
        };
    }
});
