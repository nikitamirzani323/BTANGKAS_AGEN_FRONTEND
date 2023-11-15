<script>
    import { Input } from "sveltestrap";
    
    import Panel from "../../components/Panel.svelte";
    import Loader from "../../components/Loader.svelte";
	import Button from "../../components/Button.svelte";
	import Modal from "../../components/Modal.svelte";
    import { createEventDispatcher } from "svelte";

    
	export let table_header_font = ""
	export let table_body_font = ""
	export let token = ""
	export let listHome = []
	export let totalrecord = 0
    let dispatch = createEventDispatcher();
	let title_page = "TRANSAKSI"
    let sData = "";
    let myModal_newentry = "";
    let flag_btnsave = true;
    let minbet_field = 0;
    let create_field = "";
    let update_field = "";
    let idrecord = 0;
 

      //COMPANY CONF POINT
    let listdetail_transaksi = [];

    let invoice_username = "";
    let invoice_roundbet = 0;
    let invoice_totalbet = 0;
    let invoice_totalwin = 0;
    let invoice_totalbonus = 0;
    let invoice_winlose = 0;
    let invoice_winlose_css = "";
    let invoice_card_result = "";
    let invoice_card_pattern_sort  = "";
    let invoice_card_win  = "";
    let invoice_card_poin  = "";

    let title_invoice = "";
    let searchHome = "";
    let filterHome = [];
    let css_loader = "display: none;";
    let msgloader = "";

    $: {
        if (searchHome) {
            filterHome = listHome.filter(
                (item) =>
                    item.home_id
                        .toLowerCase()
                        .includes(searchHome.toLowerCase())
            );
        } else {
            filterHome = [...listHome];
        }
    }
    
    const NewData = (e,id,card_result,card_win,card_poin,username,roundbet,totalbet,totalwin,totalbonus,winlose,winlose_css) => {
        sData = e
        if(sData == "New"){
            
        }else{
            // flag_code = true;
            idrecord = id
            title_invoice = id;
            invoice_card_result = card_result;
            invoice_card_win = card_win;
            invoice_card_poin = card_poin;
            invoice_username = username;
            invoice_roundbet = roundbet;
            invoice_totalbet = totalbet;
            invoice_totalwin = totalwin;
            invoice_totalbonus = totalbonus;
            invoice_winlose = winlose;
            invoice_winlose_css = winlose_css;


            let shuffleArray = []
            let temp_data = invoice_card_result.split('-')
            let temp_data_total = temp_data.length
            invoice_card_pattern_sort = ""
            for(let i=0;i<temp_data_total;i++) {
                shuffleArray.push(card_result_data[temp_data[i]])
            }
        
            function compareByval_display(a, b) {
                return a.val_display - b.val_display;
            }
            let pattern_array_sort = shuffleArray.sort(compareByval_display);
            for(let i=0;i<pattern_array_sort.length;i++) {
            
                if(i == pattern_array_sort.length-1) {
                    let temp_data = card_result_data.findIndex(card => card.id ==pattern_array_sort[i].id)
                    invoice_card_pattern_sort += temp_data
                }else{
                    let temp_data = card_result_data.findIndex(card => card.id ==pattern_array_sort[i].id)
                    invoice_card_pattern_sort += temp_data+"-"
                }
            }


            call_detail(id) 
            myModal_newentry = new bootstrap.Modal(document.getElementById("modaldetailtransaksi"));
            myModal_newentry.show();
        }
       
        
    };
    
    
    async function call_detail(idinvoice) {
        listdetail_transaksi = [];
        const res = await fetch("/api/transaksidetail", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                invoice: idinvoice,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                let no = 0;
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    listdetail_transaksi = [
                        ...listdetail_transaksi,
                        {
                        detail_no: no,
                        detail_id: record[i]["transaksidetail_id"],
                        detail_date: record[i]["transaksidetail_date"],
                        detail_round: record[i]["transaksidetail_roundbet"],
                        detail_bet: record[i]["transaksidetail_bet"],
                        detail_win: record[i]["transaksidetail_win"],
                        detail_bonus: record[i]["transaksidetail_bonus"],
                        detail_card_codepoin: record[i]["transaksidetail_card_codepoin"],
                        detail_card_win: record[i]["transaksidetail_card_win"],
                        detail_creditbefore: record[i]["transaksidetail_creditbefore"],
                        detail_creditafter: record[i]["transaksidetail_creditafter"],
                        detail_status: record[i]["transaksidetail_status"],
                        detail_status_css: record[i]["transaksidetail_status_css"],
                        detail_create: record[i]["transaksidetail_create"],
                        detail_update: record[i]["transaksidetail_update"],
                        },
                    ];
                }
            }
        }
    }
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    
    
   
    
    
    function callFunction(event){
        switch(event.detail){
            case "NEW":
                NewData("New","","","");
                break;
            case "REFRESH":
                RefreshHalaman();break;
            case "SAVE":
                handleSubmit();break;
        }
    }
    const handleKeyboard_checkenter = (e) => {
        let keyCode = e.which || e.keyCode;
        if (keyCode === 13) {
                filterTafsirMimpi = [];
                listHome = [];
                const tafsir = {
                    searchTafsirMimpi,
                };
                dispatch("handleTafsirMimpi", tafsir);
        }  
    };
 
    const card_result_data = [
      {id:"2_diamond",val:"2",val_display:2,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_2.png"},
      {id:"3_diamond",val:"3",val_display:3,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_3.png"},
      {id:"4_diamond",val:"4",val_display:4,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_4.png"},
      {id:"5_diamond",val:"5",val_display:5,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_5.png"},
      {id:"6_diamond",val:"6",val_display:6,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_6.png"},
      {id:"7_diamond",val:"7",val_display:7,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_7.png"},
      {id:"8_diamond",val:"8",val_display:8,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_8.png"},
      {id:"9_diamond",val:"9",val_display:9,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_9.png"},
      {id:"10_diamond",val:"10",val_display:10,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_10.png"},
      {id:"j_diamond",val:"J",val_display:11,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_J.png"},
      {id:"q_diamond",val:"Q",val_display:12,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_Q.png"},
      {id:"k_diamond",val:"K",val_display:13,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_K.png"},
      {id:"as_diamond",val:"AS",val_display:14,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_AS.png"},
      {id:"2_heart",val:"2",val_display:2,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_2.png"},
      {id:"3_heart",val:"3",val_display:3,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_3.png"},
      {id:"4_heart",val:"4",val_display:4,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_4.png"},
      {id:"5_heart",val:"5",val_display:5,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_5.png"},
      {id:"6_heart",val:"6",val_display:6,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_6.png"},
      {id:"7_heart",val:"7",val_display:7,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_7.png"},
      {id:"8_heart",val:"8",val_display:8,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_8.png"},
      {id:"9_heart",val:"9",val_display:9,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_9.png"},
      {id:"10_heart",val:"10",val_display:10,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_10.png"},
      {id:"j_heart",val:"J",val_display:11,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_J.png"},
      {id:"q_heart",val:"Q",val_display:12,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_Q.png"},
      {id:"k_heart",val:"K",val_display:13,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_K.png"},
      {id:"as_heart",val:"AS",val_display:14,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_AS.png"},
      {id:"2_club",val:"2",val_display:2,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_2.png"},
      {id:"3_club",val:"3",val_display:3,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_3.png"},
      {id:"4_club",val:"4",val_display:4,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_4.png"},
      {id:"5_club",val:"5",val_display:5,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_5.png"},
      {id:"6_club",val:"6",val_display:6,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_6.png"},
      {id:"7_club",val:"7",val_display:7,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_7.png"},
      {id:"8_club",val:"8",val_display:8,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_8.png"},
      {id:"9_club",val:"9",val_display:9,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_9.png"},
      {id:"10_club",val:"10",val_display:10,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_10.png"},
      {id:"j_club",val:"J",val_display:11,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_J.png"},
      {id:"q_club",val:"Q",val_display:12,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_Q.png"},
      {id:"k_club",val:"K",val_display:13,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_K.png"},
      {id:"as_club",val:"AS",val_display:14,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_AS.png"},
      {id:"2_spade",val:"2",val_display:2,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_2.png"},
      {id:"3_spade",val:"3",val_display:3,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_3.png"},
      {id:"4_spade",val:"4",val_display:4,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_4.png"},
      {id:"5_spade",val:"5",val_display:5,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_5.png"},
      {id:"6_spade",val:"6",val_display:6,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_6.png"},
      {id:"7_spade",val:"7",val_display:7,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_7.png"},
      {id:"8_spade",val:"8",val_display:8,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_8.png"},
      {id:"9_spade",val:"9",val_display:9,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_9.png"},
      {id:"10_spade",val:"10",val_display:10,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_10.png"},
      {id:"j_spade",val:"J",val_display:11,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_J.png"},
      {id:"q_spade",val:"Q",val_display:12,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_Q.png"},
      {id:"k_spade",val:"K",val_display:13,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_K.png"},
      {id:"as_spade",val:"AS",val_display:14,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_AS.png"},
      {id:"jk_black",val:"JK",val_display:1,code_card:"JK",img:"./CARD/WHITE/CARD_JOKER_BLACK.png"},
      {id:"jk_red",val:"JK",val_display:1,code_card:"JK",img:"./CARD/WHITE/CARD_JOKER_RED.png"},
    ]
    function card_img(e,y){ 
      if(e != "" || e.length > 0){
        let data = e.split("-");
        let total_data = e.split("-").length;
        let img_data = "";
        for(let i=0;i<total_data;i++){
        //   const searchIndex = card_result_data.findIndex((car) => car.id==data[i]);
          
          img_data +="<img width='"+y+"px' src='"+card_result_data[data[i]].img+"' /> "
        }
        return img_data
      }else{
        return ""
      }
    }
    function card_img2(e,y){ 
      if(e != "" || e.length > 0){
        let data = e.split(",");
        let total_data = e.split(",").length;
        let img_data = "";
        for(let i=0;i<total_data;i++){
          const searchIndex = card_result_data.findIndex((car) => car.id==data[i]);
          
          img_data +="<img width='"+y+"px' src='"+card_result_data[searchIndex].img+"' /> "
        }
        return img_data
      }else{
        return ""
      }
    }
</script>
<div id="loader" style="margin-left:50%;{css_loader}">
    {msgloader}
</div>
<div class="container" style="margin-top: 70px;">
    <div class="row">
        <div class="col-sm-12">
            <Button
                on:click={callFunction}
                button_function="REFRESH"
                button_title="Refresh"
                button_css="btn-primary"/>
            <Panel
                card_search={true}
                card_title="{title_page}"
                card_footer={totalrecord}>
                <slot:template slot="card-search">
                    <div class="col-lg-12" style="padding: 5px;">
                        <input
                            bind:value={searchHome}
                            on:keypress={handleKeyboard_checkenter}
                            type="text"
                            class="form-control"
                            placeholder="Search Invoice"
                            aria-label="Search"/>
                    </div>
                </slot:template>
                <slot:template slot="card-body">
                    <table class="table table-striped ">
                        <thead>
                            <tr>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;">&nbsp;</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                                <th NOWRAP width="1%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">&nbsp;</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">INVOICE</th>
                                <th NOWRAP width="10%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">DATE</th>
                                <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">USERNAME</th>
                                <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">PATTERN</th>
                                <th NOWRAP width="10%" style="text-align: right;vertical-align: top;font-weight:bold;font-size: {table_header_font};">ROUNDBET</th>
                                <th NOWRAP width="10%" style="text-align: right;vertical-align: top;font-weight:bold;font-size: {table_header_font};">TOTALBET</th>
                                <th NOWRAP width="10%" style="text-align: right;vertical-align: top;font-weight:bold;font-size: {table_header_font};">TOTALBONUS</th>
                                <th NOWRAP width="10%" style="text-align: right;vertical-align: top;font-weight:bold;font-size: {table_header_font};">TOTALWIN</th>
                                <th NOWRAP width="10%" style="text-align: right;vertical-align: top;font-weight:bold;font-size: {table_header_font};">WINLOSE</th>
                            </tr>
                        </thead>
                        {#if totalrecord > 0}
                        <tbody>
                            {#each filterHome as rec }
                                <tr>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i on:click={() => {
                                                NewData("Edit",rec.home_id,rec.home_card_result,rec.home_card_win,rec.home_card_codepoin,
                                                rec.home_username,rec.home_roundbet,rec.home_totalbet,rec.home_totalwin,rec.home_totalbonus,rec.home_winlose,rec.home_winlose_css);
                                            }} class="bi bi-pencil"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_no}</td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">
                                        <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec.home_status_css}">
                                            {rec.home_status}
                                        </span>
                                    </td>
                                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_id}</td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_date}</td>
                                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_username}</td>
                                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                                        {rec.home_card_pattern}<br />
                                        {rec.home_card_result}&nbsp;&nbsp;{rec.home_card_codepoin}
                                    </td>
                                    <td  style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;">{new Intl.NumberFormat().format(rec.home_roundbet)}</td>
                                    <td  style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;">{new Intl.NumberFormat().format(rec.home_totalbet)}</td>
                                    <td  style="text-align: right;vertical-align: top;font-size: {table_body_font};color:red;">{new Intl.NumberFormat().format(rec.home_totalbonus)}</td>
                                    <td  style="text-align: right;vertical-align: top;font-size: {table_body_font};color:red;">{new Intl.NumberFormat().format(rec.home_totalwin)}</td>
                                    <td  style="text-align: right;vertical-align: top;font-size: {table_body_font};{rec.home_winlose_css}">{new Intl.NumberFormat().format(rec.home_winlose)}</td>
                                </tr>
                            {/each}
                        </tbody>
                        {:else}
                        <tbody>
                            <tr>
                                <td colspan="20">
                                    <center>
                                        <Loader />
                                    </center>
                                </td>
                            </tr>
                        </tbody>
                        {/if} 
                    </table>
                </slot:template>
            </Panel>
        </div>
    </div>
</div>



<Modal
	modal_id="modaldetailtransaksi"
	modal_size="modal-dialog-centered modal-lg"
	modal_title="INVOICE - {title_invoice}"
    modal_body_css="height:700px; overflow-y: scroll;"
    modal_footer_css="padding:5px;"
	modal_footer={false}>
	<slot:template slot="body">
        <table style="width: 100%;">
            <tr>
                <td width="50%" style="font-size: 12px;text-align: left;">USERNAME</td>
                <td width="1%" style="font-size: 12px;text-align: left;">:</td>
                <td width="*" style="font-size: 12px;text-align: left;">{invoice_username}</td>
            </tr>
            <tr>
                <td style="font-size: 12px;text-align: left;">ROUNDBET</td>
                <td style="font-size: 12px;text-align: left;">:</td>
                <td style="font-size: 12px;text-align: right;color:blue;font-weight: bold;">{new Intl.NumberFormat().format(invoice_roundbet)}</td>
            </tr>
            <tr>
                <td style="font-size: 12px;text-align: left;">TOTAL BET</td>
                <td style="font-size: 12px;text-align: left;">:</td>
                <td style="font-size: 12px;text-align: right;color:blue;font-weight: bold;">{new Intl.NumberFormat().format(invoice_totalbet)}</td>
            </tr>
            <tr>
                <td style="font-size: 12px;text-align: left;">TOTAL WIN</td>
                <td style="font-size: 12px;text-align: left;">:</td>
                <td style="font-size: 12px;text-align: right;color:red;font-weight: bold;">{new Intl.NumberFormat().format(invoice_totalwin)}</td>
            </tr>
            <tr>
                <td style="font-size: 12px;text-align: left;">TOTAL BONUS</td>
                <td style="font-size: 12px;text-align: left;">:</td>
                <td style="font-size: 12px;text-align: right;color:red;font-weight: bold;">{new Intl.NumberFormat().format(invoice_totalbonus)}</td>
            </tr>
            <tr>
                <td style="font-size: 12px;text-align: left;">TOTAL WINLOSE</td>
                <td style="font-size: 12px;text-align: left;">:</td>
                <td style="font-size: 12px;text-align: right;{invoice_winlose_css}font-weight: bold;">
                    {new Intl.NumberFormat().format(invoice_winlose)}
                </td>
            </tr>
        </table>
        CARD RESULT<br />
        {@html card_img(invoice_card_result,85)}<br /><br />
        CARD RESULT - SORT<br />
        {@html card_img(invoice_card_pattern_sort,85)}<br /><br />
        CARD WIN : {invoice_card_poin}<br />
        {@html card_img2(invoice_card_win,85)}<br /><br />

        <table class="table table-striped ">
            <thead>
                <tr>
                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">&nbsp;</th>
                    <th NOWRAP width="*" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">DATE</th>
                    <th NOWRAP width="10%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">ROUND</th>
                    <th NOWRAP width="10%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">BET</th>
                    <th NOWRAP width="10%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">WIN</th>
                    <th NOWRAP width="10%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">BONUS</th>
                    <th NOWRAP width="15%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">C-BEFORE</th>
                    <th NOWRAP width="15%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">C-AFTER</th>
                </tr>
            </thead>
            <tbody>
                {#each listdetail_transaksi as rec }
                    <tr>
                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.detail_no}</td>
                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                            <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec.detail_status_css}">
                                {rec.detail_status}
                            </span>
                        </td>
                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.detail_date}</td>
                        <td  style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;">{new Intl.NumberFormat().format(rec.detail_round)}</td>
                        <td  style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;">{new Intl.NumberFormat().format(rec.detail_bet)}</td>
                        <td  style="text-align: right;vertical-align: top;font-size: {table_body_font};color:red;">{new Intl.NumberFormat().format(rec.detail_win)}</td>
                        <td  style="text-align: right;vertical-align: top;font-size: {table_body_font};color:red;">{new Intl.NumberFormat().format(rec.detail_bonus)}</td>
                        <td  style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;">{new Intl.NumberFormat().format(rec.detail_creditbefore)}</td>
                        <td  style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;">{new Intl.NumberFormat().format(rec.detail_creditafter)}</td>
                    </tr>
                {/each}
            </tbody>
        </table>
	</slot:template>
	<slot:template slot="footer">
        
	</slot:template>
</Modal>

