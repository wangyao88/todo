function openlayer(id){
    layer.open({
        type: 2,
        title: '修改密码',
        shadeClose: false,
        shade: 0.5,
        skin: 'layui-layer-rim',
//            maxmin: true,
        closeBtn:2,
        area: ['35%', '40%'],
        content: 'password.html'
        //iframe的url
    });
}