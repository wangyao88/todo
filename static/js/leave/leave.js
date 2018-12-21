!function(){
    laydate.skin('dahong');
    laydate({elem: '#LeaveStartTime'});
    laydate({elem: '#LeaveEndTime'});
}();

function leaveload() {
    $('#table').bootstrapTable({
        method: "get",
        url: "/leave/doList",
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
                leaveType: $("#leaveType").val(),
                leaveStartTime: $("#LeaveStartTime").val(),
                leaveEndTime: $("#LeaveEndTime").val()
            }
            return temp
        },
        columns: [
            {
                title: '编号',
                field: 'LeaveId',
                align: 'center',
                valign: 'middle',
                hidden: true
            },
            {
                title: '请假类型',
                field: 'LeaveType',
                align: 'center',
                valign: 'middle'
            },
            {
                title: '请假事由',
                field: 'LeaveNote',
                align: 'center'
            },
            {
                title: '创建时间',
                field: 'LeaveCreateTime',
                align: 'center',
                formatter: function (value, row) {
                    return row.LeaveCreateTime.substr(0,10);
                }
            },
            {
                title: '开始时间',
                field: 'LeaveStartTime',
                align: 'center',
                formatter: function (value, row) {
                    return row.LeaveStartTime.substr(0,10);
                }
            },
            {
                title: '结束时间',
                field: 'LeaveEndTime',
                align: 'center',
                formatter: function (value, row) {
                    if(row.LeaveEndTime == "0001-01-01T00:00:00Z") {
                        return "";
                    }
                    var subTime = row.LeaveEndTime.substr(0,10);
                    if(subTime == "0001-01-01") {
                        return "";
                    }
                    return subTime
                }
            },
            {
                title: '请假天数',
                field: 'LeaveSum',
                align: 'center'
            },
            {
                title: '操作',
                field: '',
                align: 'center',
                formatter: function (value, row) {
                    var del = '<button button="#" mce_href="#" onclick="delLeave(\'' + row.LeaveId + '\')">删除</button> ';
                    var edit = '<button button="#" mce_href="#" onclick="editLeave(\'' + row.LeaveId + '\')">编辑</button> ';
                    return del + edit;
                }
            }
        ]
    });
}

function getLeaveTableData() {
    $('#table').bootstrapTable('refresh');
}

function addLeave() {
    var url = '/leave/add'
    openlayer(url)
}

function editLeave(LeaveId) {
    var url = '/leave/update?LeaveId='+LeaveId
    openlayer(url)
}

function detailLeave(LeaveId) {
    var url = '/leave/detail?LeaveId='+LeaveId
    openlayer(url)
}

function delLeave(LeaveId) {
    $.ajax({
        url: '/leave/delete',
        type: 'post',
        data: {
            LeaveId: LeaveId
        },
        dataType: 'json',
        success: function (data) {
            if (data.status == 1) {
                alert("删除成功！")
                getLeaveTableData();
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
        title: '请假信息',
        shadeClose: true,
        shade: 0.5,
        skin: 'layui-layer-rim',
        area: ['98%', '90%'],
        shadeClose: true,
        closeBtn: 2,
        content: url
    });
}





