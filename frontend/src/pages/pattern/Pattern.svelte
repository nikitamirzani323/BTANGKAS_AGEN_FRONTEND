<script>
    import Home from "./Home.svelte";
   
    
    export let table_header_font = "";
	export let table_body_font = "";
    
    let token = localStorage.getItem("token");
    let akses_page = false;
    let listHome = [];
    let listPoin = [];
    let listPage = [];
    let search = "";
    let search_status = "";
    let record = "";
    let record_message = "";
    let perpage = 0;
    let page = 0;
    let totalrecordall = 0;
    let totalpaging = 0;
    let totalrecord = 0;
    let totallose = 0;
    let totalwin = 0;

    async function initapp() {
        const res = await fetch("/api/valid", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                page: "CURR-VIEW",
            }),
        });
        const json = await res.json();
        if (json.status === 400) {
            logout();
        } else if (json.status == 403) {
            alert(json.message);
        } else {
            akses_page = true;
            initHome();
        }
    }
    async function initHome(e,y) {
        listHome = [];
        const res = await fetch("/api/pattern", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                pattern_search: e,
                pattern_search_status:y,
                pattern_page : parseInt(page)
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            record = json.record;
            let record_listpoint = json.listpoint;
            perpage = json.perpage;
            totallose = json.totallose;
            totalwin = json.totalwin;
            totalrecordall = json.totalrecord;
            record_message = json.message;
            if (record != null) {
                totalpaging = Math.ceil(parseInt(totalrecordall) / parseInt(perpage))
                totalrecord = totalrecordall;
                let no = 0
                if(page > 1){
                    no = parseInt(page) 
                }
                for (var i = 0; i < record.length; i++) {
                    no = parseInt(no) + 1;
                    listHome = [
                        ...listHome,
                        {
                            home_no: no,
                            home_id: record[i]["pattern_id"],
                            home_card: record[i]["pattern_idcard"],
                            home_codepoin: record[i]["pattern_codepoin"],
                            home_nmpoin: record[i]["pattern_nmpoin"],
                            home_resultcardwin: record[i]["pattern_resultcardwin"],
                            home_status: record[i]["pattern_status"],
                            home_status_css: record[i]["pattern_status_css"],
                            home_create: record[i]["pattern_create"],
                            home_update: record[i]["pattern_update"],
                        },
                    ];
                }
                listPage = [];
                for(var i=1;i<=totalpaging;i++){
                    listPage = [
                        ...listPage,
                        {
                            page_id: i,
                            page_value: ((i*perpage)-perpage),
                            page_display: i + " Of " + perpage*i,
                        },
                    ];
                }
            }
            for (var i = 0; i < record_listpoint.length; i++) {
                listPoin = [
                    ...listPoin,
                    {
                        patternlistpoint_id: record_listpoint[i]["patternlistpoint_id"],
                        patternlistpoint_codepoin: record_listpoint[i]["patternlistpoint_codepoin"],
                        patternlistpoint_nmpoin: record_listpoint[i]["patternlistpoint_nmpoin"],
                        patternlistpoint_total: record_listpoint[i]["patternlistpoint_total"],
                    },
                ];
            }
        } else {
            logout();
        }
    }
    async function logout() {
        localStorage.clear();
        window.location.href = "/";
    }
    const handleRefreshData = (e) => {
        listHome = [];
        totalrecord = 0;
        setTimeout(function () {
            initHome();
        }, 500);
    };
    const handlePaging = (e) => {
        page = e.detail.page
        search_status = e.detail.status_search
        initHome("",search_status)
    };
    const handleSearch = (e) => {
        search = e.detail.searchHome;
        initHome(search,"")
    };
    const handleSearchStatus = (e) => {
        search_status = e.detail.searchstatus;
        initHome("",search_status)
    };
    initapp()
</script>

{#if akses_page == true}
<Home
    on:handlePaging={handlePaging}
    on:handleSearch={handleSearch}
    on:handleSearchStatus={handleSearchStatus}
    on:handleRefreshData={handleRefreshData}
    {token}
    {table_header_font}
    {table_body_font}
    {listPage}
    {listHome}
    {listPoin}
    {totallose}
    {totalwin}
    {totalrecord}/>
{/if}