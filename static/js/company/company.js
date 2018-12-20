!function(){
    laydate.skin('dahong');//切换皮肤，请查看skins下面皮肤库
    laydate({elem: '#demo'});//绑定元素
}();

function companyload() {
    $('#table').bootstrapTable({
        method: "get",
        url: "/company/doList",
        striped: true,
        singleSelect: false,
        dataType: "json",
        pagination: true, //分页
        sidePagination: "server",
        pageSize: 10,
        pageNumber: 1,
        search: false, //显示搜索框
        queryParams: function(params) {
            var temp = {
                pageSize: params.limit,
                pageNo: (params.offset / params.limit) + 1,
                companyName: $("#companyName").val()
            }
            return temp
        },
        columns: [
            {
                field: 'CompanyId',
                align: 'center',
                valign: 'middle',
                hidden: true
            },
            {
                title: '公司名称',
                field: 'CompanyName',
                align: 'center',
                valign: 'middle'
            },
            {
                title: '担任职位',
                field: 'CompanyJob',
                align: 'center',
                valign: 'middle'
            },
            {
                title: '入职时间',
                field: 'CompanyInTime',
                align: 'center',
                formatter: function (value, row) {
                    return row.CompanyInTime.substr(0,10);
                }
            },
            {
                title: '离职时间',
                field: 'CompanyOutTime',
                align: 'center',
                formatter: function (value, row) {
                    if(row.CompanyOutTime == "0001-01-01T00:00:00Z") {
                        return "至今";
                    }
                    return row.CompanyOutTime.substr(0,10);
                }
            },
            {
                title: '公司描述',
                field: 'CompanyDescription',
                align: 'center'
            },

            {
                title: '操作',
                field: '',
                align: 'center',
                formatter: function (value, row) {
                    var e = '<button button="#" mce_href="#" onclick="delCompany(\'' + row.CompanyId + '\')">删除</button> '
                    var d = '<button button="#" mce_href="#" onclick="editCompany(\'' + row.CompanyId + '\')">编辑</button> ';
                    return e + d;
                }
            }
        ]
    });
}

function getCompanyTableData() {
    $('#table').bootstrapTable('refresh');
}

function addCompany() {
    var url = '/company/add'
    openlayer(url)
}

function editCompany(CompanyId) {
    var url = '/company/update?CompanyId='+CompanyId
    openlayer(url)
}

function delCompany(CompanyId) {
    $.ajax({
        url: '/company/delete',
        type: 'post',
        data: {
            CompanyId: CompanyId
        },
        dataType: 'json',
        success: function (data) {
            if (data.status == 1) {
                alert("删除成功！")
                getCompanyTableData();
            } else {
                alert("删除失败")
            }
        },
        error: function (err) {
        }
    });
}

function openlayer(url) {
    layer.open({
        type: 2,
        title: '公司信息',
        shadeClose: true,
        shade: 0.5,
        skin: 'layui-layer-rim',
        area: ['98%', '90%'],
        shadeClose: true,
        closeBtn: 2,
        content: url
    });
}





