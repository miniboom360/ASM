import { computed, inject, provide } from 'vue';
export function useSize(props) {
    // 组件尺寸上下文
    var xesize = inject('xesize', null);
    var computeSize = computed(function () {
        return props.size || (xesize ? xesize.value : null);
    });
    provide('xesize', computeSize);
    return computeSize;
}
