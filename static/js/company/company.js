!function(){
    laydate.skin('dahong');//切换皮肤，请查看skins下面皮肤库
    laydate({elem: '#demo'});//绑定元素
}();

var noticeTitle, Publisher, currentID, noticeTime, flag = true;

function companyload() {
    $('#table').bootstrapTable({
        method: "get",
        striped: true,
        singleSelect: false,
        dataType: "json",
        pagination: true, //分页
        pageSize: 10,
        pageNumber: 1,
        search: false, //显示搜索框
        contentType: "application/x-www-form-urlencoded",
        queryParams: null,
        columns: [

            {
                title: "编号",
                field: 'ID',
                align: 'center',
                valign: 'middle'
            },
            {
                title: '标题',
                field: 'TITLE',
                align: 'center',
                valign: 'middle'
            },
            {
                title: '发布人',
                field: 'PUBLISHER',
                align: 'center',
                valign: 'middle'
            },
            {
                title: '发布时间',
                field: 'PUBTIME',
                align: 'center'
            },
            {
                title: '发布内容',
                field: 'PUBCONENT',
                align: 'center'
            },

            {
                title: '操作',
                field: '',
                align: 'center',
                formatter: function (value, row) {
                    var e = '<button button="#" mce_href="#" onclick="delNotice(\'' + row.WORKRECORDID + '\')">删除</button> '
                    var d = '<button button="#" mce_href="#" onclick="editNotice(\'' + row.WORKRECORDID + '\')">编辑</button> ';
                    return e + d;
                }
            }
        ]
    });
    getCompanyTableData();
}
function getCompanyTableData() {
    if (flag) {
        noticeTitle = "";
        Publisher = "";
        noticeTime = "";
        flag = false;
    } else {
        noticeTitle = $("#noticeTit").val();
        Publisher = $("#noticePuer").val();
        noticeTime = $("#demo").val();
    }
    $.ajax({
        type: "GET",
        url: "../WorkRecord/SearchWork?dtStart=" + noticeTitle + "&dtEnd=" + Publisher + "&dtEnd=" + noticeTime,
        dataType: "json",
        success: function (result) {
            if (result.data) {
                var NoticeTableData = result.data;
                $('#table').bootstrapTable("load", NoticeTableData);
            }
        }
    })
}
function addCompany() {
    openlayer()
    currentID = "";
}
function editCompany(id) {
    openlayer()
    currentID = id;
}
function delCompany(id) {
    var NoticeId = id;
    $.ajax({
        url: '../WorkRecord/DeleteWork?workId=' + NoticeId,
        type: 'post',
        dataType: 'json',
        success: function (data) {
            if (data.status == 1) {
                alert("删除成功！")
                getNoticeTableData();
            } else {
                alert("删除失败")
            }
        },
        error: function (err) {
        }
    });
}
function getCurrentID() {
    return currentID;
}
function openlayer() {
    layer.open({
        type: 2,
        title: '公司信息',
        shadeClose: true,
        shade: 0.5,
        skin: 'layui-layer-rim',
        area: ['98%', '98%'],
        shadeClose: true,
        closeBtn: 2,
        content: '/company/add'
    });

}





