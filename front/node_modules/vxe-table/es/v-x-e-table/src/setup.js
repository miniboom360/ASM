import GlobalConfig from './conf';
import XEUtils from 'xe-utils';
/**
 * 全局参数设置
 */
export var setup = function (options) {
    return XEUtils.merge(GlobalConfig, options);
};
