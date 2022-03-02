import baiduAnalytics from 'vue-baidu-analytics'

export default ({ Vue, router }) => {
  Vue.use(baiduAnalytics, {
    router: router,
    siteIdList: [
      '45951f610a1fa82985715b79291a8de9',
    ],
    isDebug: false
  });
};