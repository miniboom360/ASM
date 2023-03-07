import { defineComponent, h, createApp, resolveComponent, reactive } from 'vue';
var dynamicContainerElem;
export var dynamicStore = reactive({
    modals: []
});
/**
 * 动态组件
 */
var VxeDynamics = defineComponent({
    setup: function () {
        return function () {
            var modals = dynamicStore.modals;
            return h('div', {
                class: 'vxe-dynamics--modal'
            }, modals.map(function (item) { return h(resolveComponent('vxe-modal'), item); }));
        };
    }
});
export var dynamicApp = createApp(VxeDynamics);
export function checkDynamic() {
    if (!dynamicContainerElem) {
        dynamicContainerElem = document.createElement('div');
        dynamicContainerElem.className = 'vxe-dynamics';
        document.body.appendChild(dynamicContainerElem);
        dynamicApp.mount(dynamicContainerElem);
    }
}
