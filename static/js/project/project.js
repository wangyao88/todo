!function(){
    laydate.skin('dahong');
    laydate({elem: '#demo'});
}();

function projectload() {
    $('#table').bootstrapTable({
        method: "get",
        url: "/project/doList",
        striped: true,
        singleSelect: false,
        dataType: "json",
        pagination: true,
        sidePagination: "server",
        pageSize: 10,
        pageNumber: 1,
        search: false,
        queryParams: function(params) {
            var temp = {
                pageSize: params.limit,
                pageNo: (params.offset / params.limit) + 1,
                projectName: $("#projectName").val()
            }
            return temp
        },
        columns: [
            {
                field: 'ProjectId',
                align: 'center',
                valign: 'middle',
                hidden: true
            },
            {
                title: '项目名称',
                field: 'ProjectName',
                align: 'center',
                valign: 'middle'
            },
            {
                title: '公司名称',
                field: 'Company.CompanyName',
                align: 'center'
            },
            {
                title: '担任职位',
                field: 'ProjectJob',
                align: 'center',
                valign: 'middle'
            },
            {
                title: '开始时间',
                field: 'ProjectStartTime',
                align: 'center',
                formatter: function (value, row) {
                    return row.ProjectStartTime.substr(0,10);
                }
            },
            {
                title: '结束时间',
                field: 'ProjectEndTime',
                align: 'center',
                formatter: function (value, row) {
                    if(row.ProjectEndTime == "0001-01-01T00:00:00Z") {
                        return "至今";
                    }
                    return row.ProjectEndTime.substr(0,10);
                }
            },
            {
                title: '项目描述',
                field: 'ProjectDescription',
                align: 'center'
            },
            {
                title: '操作',
                field: '',
                align: 'center',
                formatter: function (value, row) {
                    var e = '<button button="#" mce_href="#" onclick="delProject(\'' + row.ProjectId + '\')">删除</button> '
                    var d = '<button button="#" mce_href="#" onclick="editProject(\'' + row.ProjectId + '\')">编辑</button> ';
                    return e + d;
                }
            }
        ]
    });
}

function getProjectTableData() {
    $('#table').bootstrapTable('refresh');
}

function addProject() {
    var url = '/project/add'
    openlayer(url)
}

function editProject(ProjectId) {
    var url = '/project/update?ProjectId='+ProjectId
    openlayer(url)
}

function delProject(ProjectId) {
    $.ajax({
        url: '/project/delete',
        type: 'post',
        data: {
            ProjectId: ProjectId
        },
        dataType: 'json',
        success: function (data) {
            if (data.status == 1) {
                alert("删除成功！")
                getProjectTableData();
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
        title: '项目信息',
        shadeClose: true,
        shade: 0.5,
        skin: 'layui-layer-rim',
        area: ['98%', '90%'],
        shadeClose: true,
        closeBtn: 2,
        content: url
    });
}





