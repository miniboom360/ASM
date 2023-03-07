import { defineComponent, h, onUnmounted, inject, ref, onMounted } from 'vue';
import { createOption, watchOption, destroyOption, assemOption } from './util';
export default defineComponent({
    name: 'VxeOption',
    props: {
        value: null,
        label: { type: [String, Number, Boolean], default: '' },
        visible: { type: Boolean, default: null },
        className: [String, Function],
        disabled: Boolean
    },
    setup: function (props, _a) {
        var slots = _a.slots;
        var elem = ref();
        var $xeselect = inject('$xeselect', {});
        var optgroup = inject('xeoptgroup', null);
        var option = createOption($xeselect, props);
        option.slots = slots;
        watchOption(props, option);
        onMounted(function () {
            assemOption($xeselect, elem.value, option, optgroup);
        });
        onUnmounted(function () {
            destroyOption($xeselect, option);
        });
        return function () {
            return h('div', {
                ref: elem
            });
        };
    }
});
