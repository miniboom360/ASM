import { defineComponent, h, onUnmounted, provide, inject, ref, onMounted } from 'vue';
import { columnProps } from './column';
import { watchColumn, assemColumn, destroyColumn } from '../../table/src/util';
import Cell from '../../table/src/cell';
export default defineComponent({
    name: 'VxeColgroup',
    props: columnProps,
    setup: function (props, _a) {
        var slots = _a.slots;
        var refElem = ref();
        var $xetable = inject('$xetable', {});
        var colgroup = inject('xecolgroup', null);
        var column = Cell.createColumn($xetable, props);
        var columnSlots = {};
        if (slots.header) {
            columnSlots.header = slots.header;
        }
        var xecolumn = { column: column };
        column.slots = columnSlots;
        column.children = [];
        provide('xecolgroup', xecolumn);
        provide('$xegrid', null);
        watchColumn(props, column);
        onMounted(function () {
            assemColumn($xetable, refElem.value, column, colgroup);
        });
        onUnmounted(function () {
            destroyColumn($xetable, column);
        });
        var renderVN = function () {
            return h('div', {
                ref: refElem
            }, slots.default ? slots.default() : []);
        };
        return renderVN;
    }
});
