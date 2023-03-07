var getAllColumns = function (columns, parentColumn) {
    var result = [];
    columns.forEach(function (column) {
        column.parentId = parentColumn ? parentColumn.id : null;
        if (column.visible) {
            if (column.children && column.children.length && column.children.some(function (column) { return column.visible; })) {
                result.push(column);
                result.push.apply(result, getAllColumns(column.children, column));
            }
            else {
                result.push(column);
            }
        }
    });
    return result;
};
export var convertToRows = function (originColumns) {
    var maxLevel = 1;
    var traverse = function (column, parent) {
        if (parent) {
            column.level = parent.level + 1;
            if (maxLevel < column.level) {
                maxLevel = column.level;
            }
        }
        if (column.children && column.children.length && column.children.some(function (column) { return column.visible; })) {
            var colSpan_1 = 0;
            column.children.forEach(function (subColumn) {
                if (subColumn.visible) {
                    traverse(subColumn, column);
                    colSpan_1 += subColumn.colSpan;
                }
            });
            column.colSpan = colSpan_1;
        }
        else {
            column.colSpan = 1;
        }
    };
    originColumns.forEach(function (column) {
        column.level = 1;
        traverse(column);
    });
    var rows = [];
    for (var i = 0; i < maxLevel; i++) {
        rows.push([]);
    }
    var allColumns = getAllColumns(originColumns);
    allColumns.forEach(function (column) {
        if (column.children && column.children.length && column.children.some(function (column) { return column.visible; })) {
            column.rowSpan = 1;
        }
        else {
            column.rowSpan = maxLevel - column.level + 1;
        }
        rows[column.level - 1].push(column);
    });
    return rows;
};
