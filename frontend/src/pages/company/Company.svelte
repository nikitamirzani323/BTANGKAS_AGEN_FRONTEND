<script>
    import Home from "./Home.svelte";
   
    
    export let table_header_font = "";
	export let table_body_font = "";
    
    let token = localStorage.getItem("token");
    let akses_page = false;
    let listHome = [];
    let listCurr = [];
    let record = "";
    let record_message = "";
    let totalrecord = 0;

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
    async function initHome() {
        const res = await fetch("/api/company", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            record = json.record;
            let record_curr = json.listcurr;
            record_message = json.message;
            if (record != null) {
                totalrecord = record.length;
                let no = 0
                let domain_css = "";
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    listHome = [
                        ...listHome,
                        {
                            home_no: no,
                            home_id: record[i]["company_id"],
                            home_name: record[i]["company_name"],
                            home_idcurr: record[i]["company_idcurr"],
                            home_start: record[i]["company_startjoin"],
                            home_end: record[i]["company_endjoin"],
                            home_nmowner: record[i]["company_nmowner"],
                            home_phoneowner: record[i]["company_phoneowner"],
                            home_emailowner: record[i]["company_emailowner"],
                            home_url: record[i]["company_url"],
                            home_status: record[i]["company_status"],
                            home_status_css: record[i]["company_status_css"],
                            home_create: record[i]["company_create"],
                            home_update: record[i]["company_update"],
                        },
                    ];
                }
            }
            for (var i = 0; i < record_curr.length; i++) {
                listCurr = [
                    ...listCurr,
                    {
                        curr_id: record_curr[i]["curr_id"],
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
        listCurr = [];
        totalrecord = 0;
        setTimeout(function () {
            initHome();
        }, 500);
    };
    initapp()
</script>

{#if akses_page == true}
<Home
    on:handleRefreshData={handleRefreshData}
    {token}
    {table_header_font}
    {table_body_font}
    {listHome}
    {listCurr}
    {totalrecord}/>
{/if}