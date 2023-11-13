<script>
    import Home from "./Home.svelte";
   
    
    export let table_header_font = "";
	export let table_body_font = "";
    
    let token = localStorage.getItem("token");
    let akses_page = false;
    let listHome = [];
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
        const res = await fetch("/api/transaksi", {
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
            record_message = json.message;
            if (record != null) {
                totalrecord = record.length;
                let no = 0
                let total_winlose = 0;
                let total_winlose_css = "";
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    total_winlose = record[i]["transaksi_totalbet"] - (record[i]["transaksi_totalbonus"] + record[i]["transaksi_totalwin"])
                    if(total_winlose > 0){
                        total_winlose_css = "color:blue;"
                    }else{
                        total_winlose_css = "color:red;"
                    }
                    listHome = [
                        ...listHome,
                        {
                            home_no: no,
                            home_id: record[i]["transaksi_id"],
                            home_date: record[i]["transaksi_date"],
                            home_username: record[i]["transaksi_username"],
                            home_roundbet: record[i]["transaksi_roundbet"],
                            home_totalbet: record[i]["transaksi_totalbet"],
                            home_totalbonus: record[i]["transaksi_totalbonus"],
                            home_totalwin: record[i]["transaksi_totalwin"],
                            home_winlose: total_winlose,
                            home_winlose_css: total_winlose_css,
                            home_card_codepoin: record[i]["transaksi_card_codepoin"],
                            home_card_pattern: record[i]["transaksi_card_pattern"],
                            home_card_win: record[i]["transaksi_card_win"],
                            home_create: record[i]["lisbet_create"],
                            home_update: record[i]["lisbet_update"],
                        },
                    ];
                }
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
    initapp()
</script>

{#if akses_page == true}
<Home
    on:handleRefreshData={handleRefreshData}
    {token}
    {table_header_font}
    {table_body_font}
    {listHome}
    {totalrecord}/>
{/if}