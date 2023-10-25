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
  export let listPage = [];
  export let listPoin = [];
	export let totalrecord = 0
	export let totallose = 0
	export let totalwin = 0
  let dispatch = createEventDispatcher();
	let title_page = "PATTERN"
  let listPatternByCode = []
  let listPagePatternByCode = []
  let selectPagingPatternByCode = "0"
  let poin_code = ""
  let poin_name = ""
  let pagePatternByCode = 0
  let perpagePatternByCode = 0
  let totalrecordallPatternByCode = 0
    let sData = "";
    let myModal_newentry = "";
    let flag_btnsave = true;
    let flag_btnsavemanual = true;
    let idrecord = "";
    let resultcard = "";
    let resultcodepoint = "";
    let resultnmpoint = "";
    let resultcodepoint_after = "";
    let resultnmpoint_after = "";
    let resultcardwin = "";
    let resultstatus = "";
    let create_field = "";
    let update_field = "";
    let searchHome = "";
    let status_search = "";
    let filterHome = [];
    let css_loader = "display: none;";
    let msgloader = "";
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
    let list_point = [
      {id:"1",code:"RF",name:"Royal Flush",poin:500},
      {id:"2",code:"5K",name:"5 Of A Kind",poin:200},
      {id:"3",code:"SF",name:"Straight Flush",poin:120},
      {id:"4",code:"4K",name:"4 Of A Kind",poin:50},
      {id:"5",code:"FH",name:"Full House",poin:7},
      {id:"6",code:"FL",name:"Flush",poin:5},
      {id:"7",code:"ST",name:"Straight",poin:3},
      {id:"8",code:"3K",name:"3 Of A Kind",poin:2},
      {id:"9",code:"2P",name:"2 Pair (10 PAIR)",poin:1},
      {id:"10",code:"AP",name:"Ace Pair",poin:1},
    ]
    const pattern_stright_1 = [2,3,4,5,6]
    const pattern_stright_2 = [3,4,5,6,7]
    const pattern_stright_3 = [4,5,6,7,8]
    const pattern_stright_4 = [5,6,7,8,9]
    const pattern_stright_5 = [6,7,8,9,10]
    const pattern_stright_6 = [7,8,9,10,11]
    const pattern_stright_7 = [8,9,10,11,12]
    const pattern_stright_8 = [9,10,11,12,13]
    const pattern_stright_9 = [10,11,12,13,14]
    const pattern_stright_10 = [14,2,3,4,5]

    const pattern_4_1 = [2,3,4,5]
    const pattern_4_2 = [3,4,5,6]
    const pattern_4_3 = [4,5,6,7]
    const pattern_4_4 = [5,6,7,8]
    const pattern_4_5 = [6,7,8,9]
    const pattern_4_6 = [7,8,9,10]
    const pattern_4_7 = [8,9,10,11]
    const pattern_4_8 = [9,10,11,12]
    const pattern_4_9 = [10,11,12,13]
    const pattern_4_10 = [11,12,13,14]
    const pattern_4_11 = [14,2,3,4]

    const pattern_3_1 = [2,3,4]
    const pattern_3_2 = [3,4,5]
    const pattern_3_3 = [4,5,6]
    const pattern_3_4 = [5,6,7]
    const pattern_3_5 = [6,7,8]
    const pattern_3_6 = [7,8,9]
    const pattern_3_7 = [8,9,10]
    const pattern_3_8 = [9,10,11]
    const pattern_3_9 = [10,11,12]
    const pattern_3_10 = [11,12,13]
    const pattern_3_11 = [12,13,14]
    const pattern_3_12 = [14,2,3]


    let point_royalflush =  0
    let point_5ofkind = 0
    let point_straightflush = 0
    let point_4ofkind = 0
    let point_fullhouse = 0
    let point_flush = 0
    let point_straight = 0
    let point_3ofkind = 0
    let point_2pair = 0
    let point_acepair = 0
    let info_result = "";
    let info_card = [];
    let shuffleArray_id = [];
    let shuffleArray = [];
    let pattern_list = [];
    let pattern = [];
    let pattern_string = "";
    let pattern_card_string = "";
    let usedIndexes = [];
    let pagingnow = 0;
    let pattern_field = "";
    let pattern_img = "";
    let pattern_img_clone = "";
    let pattern_codepoin = "";
    let pattern_poin = "";
    let pattern_win = "";
    let pattern_status = "";
    let pattern_status_dua = "";
    function shuffleArray_card(array){
        let i = 0
        while(i<7){
            let randomNumber = Math.floor(Math.random() * array.length)
            if(!usedIndexes.includes(randomNumber)){
                shuffleArray_id.push(array[randomNumber].id);
                shuffleArray.push(array[randomNumber]);
                pattern.push(randomNumber);
                usedIndexes.push(randomNumber);
                i++;
            }
        }
        // console.log(shuffleArray)  
        
        for(let j = 0; j < pattern.length; j++){
            if(j==6){
                pattern_string += pattern[j]
            }else{
                pattern_string += pattern[j]+"-"
            }
        }
        for(let j = 0; j < shuffleArray_id.length; j++){
            if(j==6){
                pattern_card_string += shuffleArray_id[j]
            }else{
                pattern_card_string += shuffleArray_id[j]+","
            }
        }
        let status = hitung_statuswinlose(shuffleArray)
       
        // console.log(status)
        let temp_status = status[0]==false?"N":"Y"
        let temp_poin = "";
        let temp_listwin = "";
        if(temp_status == "Y"){
            temp_poin = list_point[status[2]].code
            // console.log("total length : " + status[1].length)
            for(let x=0;x<status[1].length;x++){
              if(x==status[1].length-1){
                temp_listwin += status[1][x].id
              }else{
                temp_listwin += status[1][x].id+","
              }
            }
            
        }
        
        const obj = {
          idpattern:pattern_string,
          idcard:pattern_card_string,
          point:temp_poin,
          resultwin:temp_listwin,
          status:temp_status,
      };

        pattern_list = [obj, ...pattern_list];
        
    
        pattern = [];
        pattern_string = "";
        pattern_card_string = "";
        temp_listwin = "";
        shuffleArray = [];
        shuffleArray_id = [];
    }
    function loopdata(){
        let i = 0
        pattern_list = [];
        usedIndexes = [];
       
        while(i<5){
            shuffleArray_card(card_result_data)
            i++;
        }

        // console.log(pattern_list)
    }
   
    $: {
        if (searchHome) {
            filterHome = listHome.filter(
                (item) =>
                    item.home_id
                        .toLowerCase()
                        .includes(searchHome.toLowerCase()) || 
                    item.home_status
                        .toLowerCase()
                        .includes(searchHome.toLowerCase()) || 
                    item.home_nmpoin
                        .toLowerCase()
                        .includes(searchHome.toLowerCase())     
            );
        } else {
            filterHome = [...listHome];
        }
        for(let p=0;p<listPoin.length;p++){
          switch(listPoin[p].patternlistpoint_codepoin){
            case "RF":
              point_royalflush = listPoin[p].patternlistpoint_total;
              break;
            case "5K":
              point_5ofkind = listPoin[p].patternlistpoint_total;
              break;
            case "SF":
              point_straightflush = listPoin[p].patternlistpoint_total;
              break;  
            case "4K":
              point_4ofkind = listPoin[p].patternlistpoint_total;
              break;
            case "FH":
              point_fullhouse = listPoin[p].patternlistpoint_total;
              break;
            case "FL":
              point_flush = listPoin[p].patternlistpoint_total;
              break;
            case "ST":
              point_straight = listPoin[p].patternlistpoint_total;
              break;
            case "3K":
              point_3ofkind = listPoin[p].patternlistpoint_total;
              break;  
            case "2P":
              point_2pair = listPoin[p].patternlistpoint_total;
              break;
            case "AP":
              point_acepair = listPoin[p].patternlistpoint_total;
              break;
          }
        }
    }
    
    const NewData = (e,id,resultcrd,code,nmpoin,resultwn,create,update) => {
        sData = e
        if(sData == "New"){
          flag_btnsave = true
          loopdata()
        }else{
          create_field = create;
          update_field = update;
          resultstatus = "";
          resultcodepoint_after = "";
          resultnmpoint_after = "";
          idrecord = id
          resultcard = resultcrd
          resultcodepoint = code
          resultnmpoint = nmpoin
          resultcardwin = resultwn

          if(resultwn == ""){
            let temp_data = id.split('-')
            let temp_data_total = temp_data.length
            shuffleArray = []
            for(let i=0;i<temp_data_total;i++) {
              shuffleArray.push(card_result_data[temp_data[i]])
            }
          
            let status = hitung_statuswinlose(shuffleArray)
            let temp_status = status[0]==false?"N":"Y"
            let temp_poin = "";
            let temp_listwin = "";
            if(temp_status == "Y"){
                resultcodepoint_after = list_point[status[2]].code
                resultnmpoint_after = list_point[status[2]].name
                // console.log("total length : " + status[1].length)
                for(let x=0;x<status[1].length;x++){
                  if(x==status[1].length-1){
                    temp_listwin += status[1][x].id
                  }else{
                    temp_listwin += status[1][x].id+","
                  }
                }
                
            }
            resultstatus = temp_status;
            resultcardwin = temp_listwin
            // console.log(resultstatus)
            if(resultcardwin == ""){
              flag_btnsave = false
            }else{
              flag_btnsave = true
            }
          }else{
            console.log("apa ya")
            let status = hitung_statuswinlose(shuffleArray)
            let temp_status = status[0]==false?"N":"Y"
            let temp_listwin = "";
            if(temp_status == "Y"){
                pattern_codepoin = list_point[status[2]].code
                pattern_poin = list_point[status[2]].name
                // console.log("total length : " + status[1].length)
                for(let x=0;x<status[1].length;x++){
                  if(x==status[1].length-1){
                    temp_listwin += status[1][x].id
                  }else{
                    temp_listwin += status[1][x].id+","
                  }
                }
                
            }
            console.log(status)
          }
        }
       
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalentrycrud"));
        myModal_newentry.show();
        
    };
    const NewManualData = () => {
        pattern_field = "";
        pattern_codepoin = "";
        pattern_poin = "";
        pattern_img = "";
        pattern_img_clone = "";
        pattern_win = "";
        pattern_status = "";
        pattern_status_dua = "";
        flag_btnsavemanual = false;
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalentrycrudmanual"));
        myModal_newentry.show();
        
    };
    const handleCheckWinLose = () => {
        let flag = true;
        pattern_poin = "";
        pattern_img = "";
        pattern_win = "";
        if(pattern_field == ""){
          alert("The pattern is required")
          flag = false;
        }
        if(flag){
          let temp_data = pattern_field.split('-')
          let temp_data_total = temp_data.length
          shuffleArray = []
          for(let i=0;i<temp_data_total;i++) {
            shuffleArray.push(card_result_data[temp_data[i]])
            if(i == temp_data_total-1) {
              pattern_img += card_result_data[temp_data[i]].id
              // pattern_img_clone += card_result_data[temp_data[i]].val_display
            }else{
              pattern_img += card_result_data[temp_data[i]].id+","
              // pattern_img_clone += card_result_data[temp_data[i]].val_display+"-"
            }
          }
         
          
          function compareByval_display(a, b) {
            return a.val_display - b.val_display;
          }
          let pattern_array_sort = shuffleArray.sort(compareByval_display);
          for(let i=0;i<pattern_array_sort.length;i++) {
         
            if(i == pattern_array_sort.length-1) {
              pattern_img_clone += pattern_array_sort[i].id
            }else{
              pattern_img_clone += pattern_array_sort[i].id+","
            }
          }

          let status = hitung_statuswinlose(shuffleArray)
          let temp_status = status[0]==false?"N":"Y"
          let temp_listwin = "";
          if(temp_status == "Y"){
              pattern_codepoin = list_point[status[2]].code
              pattern_poin = list_point[status[2]].name
              // console.log("total length : " + status[1].length)
              for(let x=0;x<status[1].length;x++){
                if(x==status[1].length-1){
                  temp_listwin += status[1][x].id
                }else{
                  temp_listwin += status[1][x].id+","
                }
              }
              
          }
          pattern_status_dua = temp_status
          pattern_win = temp_listwin

          flag_btnsavemanual = true;
        }
        
    };
    function check_pattern_status(e){
        let temp_data = e.split('-')
        let temp_data_total = temp_data.length
        let shuffleArrayTest = []
        for(let i=0;i<temp_data_total;i++) {
          shuffleArrayTest.push(card_result_data[temp_data[i]])
        }
    
      
        let status = hitung_statuswinlose(shuffleArrayTest)
        let temp_status = status[0]==false?"N":"Y"
        
        return temp_status
    }
    function hitung_statuswinlose(data_array){
      let data_result = [];
      
      // data_result = straight_factory(data_array);
      // console.log(data_result[0]);
      data_result = royal_flush_factory(data_array);
      if(!data_result[0]){
        data_result = five_kind_factory(data_array);
        if(!data_result[0]){
          data_result = straight_flush_factory(data_array);
          if(!data_result[0]){
            data_result = fourofkind_factory(data_array);
            if(!data_result[0]){
              data_result = fullhouse_factory(data_array);
              if(!data_result[0]){
                data_result = flush_factory(data_array);
                if(!data_result[0]){
                  data_result = straight_factory(data_array);
                  if(!data_result[0]){
                    data_result = threeofkind_factory(data_array);
                    if(!data_result[0]){
                      data_result = twopair_factory(data_array);
                      if(!data_result[0]){
                        data_result = acepair_factory(data_array);
                      }
                    }
                  }
                }
              }
            }
          }
        }
      }
      
      return [data_result[0],data_result[1],data_result[2]]
    }
    function royal_flush_factory(data_array){
      let flag_func =false
      let data_win = []
   
      let counts = []
      for(let i=0;i<data_array.length;i++){
        if(counts[data_array[i].code_card]){
          counts[data_array[i].code_card] += 1
        }else{
          counts[data_array[i].code_card] = 1
        }
      }
      let temp = [];
      for(let prop in counts){
        if (counts[prop] >= 3){
            if(prop == "S"){
              temp.push(prop + ":" + counts[prop])
            }
        }
      }
      // console.log(temp)
      if(temp.length > 0){
        let temp_string = temp[0]
        let temp_result = temp_string.split(":");
        let total_temp = temp_result[1];
        let total_jk = 0;
        let total_card = 0;
        if(temp_result[0] == "S"){
          if(parseInt(total_temp) == 5 || parseInt(total_temp) == 6){
            for(let i=0;i<data_array.length;i++){
              if(data_array[i].code_card == "S" || data_array[i].code_card == "JK"){
                switch(data_array[i].val){
                    case "10":
                      total_jk = total_jk + 1;break;
                    case "J":
                      total_jk = total_jk + 1;break;
                    case "K":
                      total_jk = total_jk + 1;break;
                    case "Q":
                      total_jk = total_jk + 1;break;
                    case "AS":
                      total_jk = total_jk + 1;break;
                    case "JK":
                      total_jk = total_jk + 1;break;
                }
              }
            }
            total_card = total_jk
          }
          if(parseInt(total_temp) == 4){
            for(let i=0;i<data_array.length;i++){
              if(data_array[i].code_card == "S" || data_array[i].code_card == "JK"){
                switch(data_array[i].val){
                    case "10":
                      total_jk = total_jk + 1;break;
                    case "J":
                      total_jk = total_jk + 1;break;
                    case "K":
                      total_jk = total_jk + 1;break;
                    case "Q":
                      total_jk = total_jk + 1;break;
                    case "AS":
                      total_jk = total_jk + 1;break;
                    case "JK":
                      total_jk = total_jk + 1;break;
                }
              }
            }
            total_card = total_jk
          }
          if(parseInt(total_temp) == 3){
            for(let i=0;i<data_array.length;i++){
              if(data_array[i].code_card == "S" || data_array[i].code_card == "JK"){
                switch(data_array[i].val){
                    case "10":
                      total_jk = total_jk + 1;break;
                    case "J":
                      total_jk = total_jk + 1;break;
                    case "K":
                      total_jk = total_jk + 1;break;
                    case "Q":
                      total_jk = total_jk + 1;break;
                    case "AS":
                      total_jk = total_jk + 1;break;
                    case "JK":
                      total_jk = total_jk + 1;break;
                }
              }
            }
            total_card = total_jk
          }
        }
       
   
        if(total_card == 5){
            info_result = "Royal Flush"
            info_card = pattern_stright_10
            flag_func = true;
          
            for(let i=0;i<data_array.length;i++){
              if(data_array[i].code_card == "S" || data_array[i].code_card == "JK"){
                  data_win.push(data_array[i])
              }
            }
          }
        

          
          

        
        
      }
      if(flag_func == false){
        data_win = [];
      }
      return [flag_func,data_win,0];
    }
    function five_kind_factory(data_array){
      let flag_func =false
      let data_win = []
      let counts = []
      for(let i=0;i<data_array.length;i++){
        if(counts[data_array[i].val_display]){
          counts[data_array[i].val_display] += 1
        }else{
          counts[data_array[i].val_display] = 1
        }
      }
      let temp = [];
      for(let prop in counts){
        if (counts[prop] >= 3){
              temp.push(prop + ":" + counts[prop])
          }
      }
      if(temp.length > 0){
        let temp_string = temp[0]
        let temp_result = temp_string.split(":");
        let total_temp = temp_result[1];
        let total_jk = 0;
        let total_card = 0;
        if(parseInt(total_temp) == 4){
          for(let i=0;i<data_array.length;i++){
            if(data_array[i].code_card == "JK"){
              total_jk = total_jk + 1;
            }
          }
          total_card = parseInt(total_temp) + total_jk
          if(total_card == 5){
            info_result = "5 Of A Kind"
            info_card = pattern_stright_10
            flag_func = true;
  
            for(let i=0;i<temp.length;i++){
              temp_string = temp[i]
              temp_result = temp_string.split(":");
              for(let i=0;i<data_array.length;i++){
                if(data_array[i].val_display == temp_result[0] || data_array[i].val == "JK"){
                  data_win.push(data_array[i])
                }
              }
            }
          }
        }
        if(parseInt(total_temp) == 3){
          for(let i=0;i<data_array.length;i++){
            if(data_array[i].code_card == "JK"){
              total_jk = total_jk + 1;
            }
          }
          total_card = parseInt(total_temp) + total_jk
          if(total_card == 5){
            info_result = "5 Of A Kind"
            info_card = pattern_stright_10
            flag_func = true;
            
           
            for(let i=0;i<temp.length;i++){
              temp_string = temp[i]
              temp_result = temp_string.split(":");
              for(let i=0;i<data_array.length;i++){
                if(data_array[i].val_display == temp_result[0] || data_array[i].val == "JK"){
                  data_win.push(data_array[i])
                }
              }
            }
          }
        }
      }
      if(flag_func == false){
        data_win = [];
      }
      return [flag_func,data_win,1];
    }
    function straight_flush_factory(data_array){
      let flag_func =false
      let data_win = []
      let counts = []
      for(let i=0;i<data_array.length;i++){
        if(counts[data_array[i].code_card]){
          counts[data_array[i].code_card] += 1
        }else{
          counts[data_array[i].code_card] = 1
        }
      }
      let temp = [];
      for(let prop in counts){
        if (counts[prop] >= 3){
              temp.push(prop + ":" + counts[prop])
          }
      }
      
      if(temp.length > 0){
        for(let z=0;z<temp.length;z++){
          let objdata_final = []
          let temp_string = temp[z]
          let temp_result = temp_string.split(":");
          if(parseInt(temp_result[1]) == 5){
            for(let i=0;i<data_array.length;i++){
                if(temp_result[0] == data_array[i].code_card){
                  objdata_final.push(data_array[i].val_display)
                }
            }
          
            let flag = []
            flag[0] = checkArray(pattern_stright_1,objdata_final)
            flag[1] = checkArray(pattern_stright_2,objdata_final)
            flag[2] = checkArray(pattern_stright_3,objdata_final)
            flag[3] = checkArray(pattern_stright_4,objdata_final)
            flag[4] = checkArray(pattern_stright_5,objdata_final)
            flag[5] = checkArray(pattern_stright_6,objdata_final)
            flag[6] = checkArray(pattern_stright_7,objdata_final)
            flag[7] = checkArray(pattern_stright_8,objdata_final)
            flag[8] = checkArray(pattern_stright_9,objdata_final)
            flag[9] = checkArray(pattern_stright_10,objdata_final)
            for(let i=0;i<flag.length;i++){
              if(flag[i] == true){
                info_result = "STRAIGHT FLUSH"
                info_card = pattern_stright_10
                flag_func = true;
                
              
                for(let i=0;i<data_array.length;i++){
                  temp_string = temp[0]
                  temp_result = temp_string.split(":");
                  if(temp_result[0] == data_array[i].code_card){
                    data_win.push(data_array[i])
                  }
                }
                break;
              }
            }
          }
          if(parseInt(temp_result[1]) == 4){
            for(let i=0;i<data_array.length;i++){
                if(temp_result[0] == data_array[i].code_card){
                  objdata_final.push(data_array[i].val_display)
                }
            }
           
            let flag = []
            flag[0] = checkArray(pattern_4_1,objdata_final)
            flag[1] = checkArray(pattern_4_2,objdata_final)
            flag[2] = checkArray(pattern_4_3,objdata_final)
            flag[3] = checkArray(pattern_4_4,objdata_final)
            flag[4] = checkArray(pattern_4_5,objdata_final)
            flag[5] = checkArray(pattern_4_6,objdata_final)
            flag[6] = checkArray(pattern_4_7,objdata_final)
            flag[7] = checkArray(pattern_4_8,objdata_final)
            flag[8] = checkArray(pattern_4_9,objdata_final)
            flag[9] = checkArray(pattern_4_10,objdata_final)
            flag[10] = checkArray(pattern_4_11,objdata_final)
            
            for(let i=0;i<flag.length;i++){
              if(flag[i] == true){
                let total_jk = 0;
                let total_card = 0;
                for(let i=0;i<data_array.length;i++){
                  if(data_array[i].code_card == "JK"){
                    total_jk = total_jk + 1
                  }
                }
                total_card = parseInt(temp_result[1]) + total_jk
                
                if(parseInt(total_card) == 5){
                  info_result = "STRAIGHT FLUSH"
                  info_card = pattern_stright_10
                  flag_func = true;
                  
                
                  for(let i=0;i<data_array.length;i++){
                    temp_string = temp[0]
                    temp_result = temp_string.split(":");
                    if(temp_result[0] == data_array[i].code_card || data_array[i].code_card == "JK"){
                      data_win.push(data_array[i])
                    }
                  }
                  break;
                }
              }
            }
          }
          if(parseInt(temp_result[1]) == 3){
            for(let i=0;i<data_array.length;i++){
                if(temp_result[0] == data_array[i].code_card){
                  objdata_final.push(data_array[i].val_display)
                }
            }
          
            let flag = []
            flag[0] = checkArray(pattern_3_1,objdata_final)
            flag[1] = checkArray(pattern_3_2,objdata_final)
            flag[2] = checkArray(pattern_3_3,objdata_final)
            flag[3] = checkArray(pattern_3_4,objdata_final)
            flag[4] = checkArray(pattern_3_5,objdata_final)
            flag[5] = checkArray(pattern_3_6,objdata_final)
            flag[6] = checkArray(pattern_3_7,objdata_final)
            flag[7] = checkArray(pattern_3_8,objdata_final)
            flag[8] = checkArray(pattern_3_9,objdata_final)
            flag[9] = checkArray(pattern_3_10,objdata_final)
            flag[10] = checkArray(pattern_3_11,objdata_final)
            flag[11] = checkArray(pattern_3_12,objdata_final)
           
            for(let i=0;i<flag.length;i++){
              if(flag[i] == true){
                let total_jk = 0;
                let total_card = 0;
                for(let i=0;i<data_array.length;i++){
                  if(data_array[i].code_card == "JK"){
                    total_jk = total_jk + 1
                  }
                }
                total_card = parseInt(temp_result[1]) + total_jk
                
                if(parseInt(total_card) == 5){
                  info_result = "STRAIGHT FLUSH"
                  info_card = pattern_stright_10
                  flag_func = true;
                  
                
                  for(let i=0;i<data_array.length;i++){
                    temp_string = temp[0]
                    temp_result = temp_string.split(":");
                    if(temp_result[0] == data_array[i].code_card || data_array[i].code_card == "JK"){
                      data_win.push(data_array[i])
                    }
                  }
                  break;
                }
              }
            }
          }
        }
      }
      if(flag_func == false){
        data_win = [];
      }
      return [flag_func,data_win,2];
    }
    function fourofkind_factory(data_array){
      let flag_func =false
      let data_win = []
      let counts = []
      for(let i=0;i<data_array.length;i++){
        if(counts[data_array[i].val]){
          counts[data_array[i].val] += 1
        }else{
          counts[data_array[i].val] = 1
        }
      }
      let temp = [];
      for(let prop in counts){
        if (counts[prop] >= 3){
              temp.push(prop + ":" + counts[prop])
          }
      }
      // console.log(temp)
      let total = 0;
      let total_temp = temp.length
      let temp_string = ""
      let temp_result;
      for(let i=0;i<total_temp;i++){
          temp_string = temp[i]
          temp_result = temp_string.split(":");
          total = total + parseInt(temp_result[1])
      }
      // console.log(total)
      if(total == 3){//FOUR OF KIND
        
        let flag_jk = false;
        for(let i=0;i<temp.length;i++){
          temp_string = temp[i]
          temp_result = temp_string.split(":");
          for(let j=0;j<data_array.length;j++){
            if(data_array[j].val == temp_result[0]){
              data_win.push(data_array[j])
            }
          }
        }
        for(let j=0;j<data_array.length;j++){
            if(data_array[j].val == "JK"){
              data_win.push(data_array[j])
              flag_jk = true
              break;
            }
        }
        if(flag_jk){
          info_result = "FOUR OF KIND"
          info_card = temp
          flag_func = true
        }
        // credit_animation(credit,3,totalbet)
      }
      if(total == 4){//FOUR OF KIND
        info_result = "FOUR OF KIND"
        info_card = temp
        flag_func = true
        
        for(let i=0;i<temp.length;i++){
          temp_string = temp[i]
          temp_result = temp_string.split(":");
          for(let j=0;j<data_array.length;j++){
            if(data_array[j].val == temp_result[0]){
              data_win.push(data_array[j])
            }
          }
        }
  
        // credit_animation(credit,3,totalbet)
      }
      if(total == 6){
        let data_baru = []
        for(let i=0;i<temp.length;i++){
          temp_string = temp[i]
          temp_result = temp_string.split(":");
          if(parseInt(temp_result[1]) == 3){
            data_baru.push(temp[i])
            
          }
        }
        for(let i=0;i<data_baru.length-1;i++){
          let temp_string2 = data_baru[i]
          let temp_result2 = temp_string2.split(":");
          
          for(let j=0;j<data_array.length;j++){
            if(data_array[j].val == temp_result2[0]){
              data_win.push(data_array[j])
            }
          }
        }
        for(let j=0;j<data_array.length;j++){
            if(data_array[j].val == "JK"){
              data_win.push(data_array[j])
            }
        }
        if(data_win.length == 4){
          info_result = "FOUR OF KIND"
          info_card = temp
          flag_func = true
          
        }
      }
      if(flag_func == false){
        data_win = [];
      }
      return [flag_func,data_win,3];
    }
    function fullhouse_factory(data_array){
      let flag_func =false
      let data_win = []
      let counts = []
      for(let i=0;i<data_array.length;i++){
        if(counts[data_array[i].val]){
          counts[data_array[i].val] += 1
        }else{
          counts[data_array[i].val] = 1
        }
      }
      let temp = [];
      for(let prop in counts){
        if (counts[prop] >= 2){
            if(prop != "JK"){
              temp.push(prop + ":" + counts[prop])
            }
          }
      }
      // console.log(temp)
      let total = 0;
      let total_all = 0;
      let total_jk = 0;
      let total_temp = temp.length
      let temp_string = ""
      let temp_result;
      for(let i=0;i<total_temp;i++){
          temp_string = temp[i]
          temp_result = temp_string.split(":");
          total = total + parseInt(temp_result[1])
      }
     
      if(total == 4){
        for(let i=0;i<temp.length;i++){
            temp_string = temp[i]
            temp_result = temp_string.split(":");
            for(let i=0;i<data_array.length;i++){
              if(data_array[i].val == temp_result[0]){
                data_win.push(data_array[i])
              }
            }
        }
        total = data_win.length
    
        for(let i=0;i<data_array.length;i++){
          if(data_array[i].val == "JK"){
            total_jk = total_jk + 1
            data_win.push(data_array[i])
          }
        }
        
        total_all = total_all + total + total_jk
        if(total_all == 6){
          let popped = data_win.pop();
          total_all = data_win.length
        }

        if(total_all == 5){
          info_result = "FULL HOUSE"
          info_card = temp
          flag_func = true
        } 
      }
      if(total == 5){//FULL HOUSE
        if(temp.length == 2){
          info_result = "FULL HOUSE"
          info_card = temp
          flag_func = true
          
          for(let i=0;i<temp.length;i++){
              temp_string = temp[i]
              temp_result = temp_string.split(":");
              for(let i=0;i<data_array.length;i++){
                if(data_array[i].val == temp_result[0]){
                  data_win.push(data_array[i])
                }
              }
          }
          // credit_animation(credit,4,totalbet)
        }
      }
      if(total == 6){
        if(temp.length == 2){
          info_result = "FULL HOUSE"
          info_card = temp
          flag_func = true
          
          for(let i=0;i<temp.length;i++){
              temp_string = temp[i]
              temp_result = temp_string.split(":");
              for(let i=0;i<data_array.length-1;i++){
                if(data_array[i].val == temp_result[0]){
                  data_win.push(data_array[i])
                }
              }
          }
        }else{
          
          let flag_jk = false;
          for(let i=0;i<temp.length-1;i++){
              temp_string = temp[i]
              temp_result = temp_string.split(":");
              for(let i=0;i<data_array.length;i++){
                if(data_array[i].val == temp_result[0]){
                  data_win.push(data_array[i])
                }
              }
          }
          // console.log(data_win.length)
          if(data_win.length < 5){
            for(let i=0;i<data_array.length;i++){
              if(data_array[i].val == "JK"){
                data_win.push(data_array[i])
                flag_jk = true;
              }
            }
          }
          if(flag_jk){
            info_result = "FULL HOUSE"
            info_card = temp
            flag_func = true
          }
        }
      }
      if(flag_func == false){
        data_win = [];
      }
      return [flag_func,data_win,4];
    }
    function flush_factory(data_array){
      let flag_func =false
      let data_win = []
      let counts = []
      for(let i=0;i<data_array.length;i++){
        if(counts[data_array[i].code_card]){
          counts[data_array[i].code_card] += 1
        }else{
          counts[data_array[i].code_card] = 1
        }
      }
      let temp = [];
      for(let prop in counts){
        if (counts[prop] >= 5){
              temp.push(prop + ":" + counts[prop])
          }
      }
      let total = 0;
      let total_temp = temp.length
      let temp_string = ""
      let temp_result;
      
      for(let i=0;i<total_temp;i++){
          temp_string = temp[i]
          temp_result = temp_string.split(":");
          
          total = total + parseInt(temp_result[1])
      }
     
      if(total_temp == 1){
        if(total == 5){ //FLUSH
          info_result = "FLUSH"
          info_card = temp
          flag_func = true
  
  
          for(let i=0;i<temp.length;i++){
            temp_string = temp[i]
            temp_result = temp_string.split(":");
            for(let i=0;i<data_array.length;i++){
              if(data_array[i].code_card == temp_result[0]){
                data_win.push(data_array[i])
              }
            }
          }
        
         
          // credit_animation(credit,5,totalbet)
        }
      }
      if(flag_func == false){
        data_win = [];
      }
      return [flag_func,data_win,5];
    }
    function straight_factory(data_array){
      let flag_func = false
      let data_win = []
      let objdata_master = []
      for(let i in data_array){
        objdata_master.push(data_array[i].val_display)
      }
      let flag = []
      
      flag[0] = checkArray(pattern_stright_1,objdata_master)
      flag[1] = checkArray(pattern_stright_2,objdata_master)
      flag[2] = checkArray(pattern_stright_3,objdata_master)
      flag[3] = checkArray(pattern_stright_4,objdata_master)
      flag[4] = checkArray(pattern_stright_5,objdata_master)
      flag[5] = checkArray(pattern_stright_6,objdata_master)
      flag[6] = checkArray(pattern_stright_7,objdata_master)
      flag[7] = checkArray(pattern_stright_8,objdata_master)
      flag[8] = checkArray(pattern_stright_9,objdata_master)
      flag[9] = checkArray(pattern_stright_10,objdata_master)

      flag[10] = checkArray(pattern_4_1,objdata_master)
      flag[11] = checkArray(pattern_4_2,objdata_master)
      flag[12] = checkArray(pattern_4_3,objdata_master)
      flag[13] = checkArray(pattern_4_4,objdata_master)
      flag[14] = checkArray(pattern_4_5,objdata_master)
      flag[15] = checkArray(pattern_4_6,objdata_master)
      flag[16] = checkArray(pattern_4_7,objdata_master)
      flag[17] = checkArray(pattern_4_8,objdata_master)
      flag[18] = checkArray(pattern_4_9,objdata_master)
      flag[19] = checkArray(pattern_4_10,objdata_master)
      flag[20] = checkArray(pattern_4_11,objdata_master)
      // console.log(objdata_master)
      // console.log(flag)
      for(let i=0;i<flag.length;i++){
        if(flag[i] == true){
          info_result = "STRAIGHT"
          info_card = pattern_stright_10
        
          switch(i){
            case 0:
              for(let t=0;t<pattern_stright_1.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_stright_1[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
            case 1:
              for(let t=0;t<pattern_stright_2.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_stright_2[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
            case 2:
              for(let t=0;t<pattern_stright_3.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_stright_3[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
            case 3:
              for(let t=0;t<pattern_stright_4.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_stright_4[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
            case 4:
              for(let t=0;t<pattern_stright_5.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_stright_5[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
            case 5:
              for(let t=0;t<pattern_stright_6.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_stright_6[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
            case 6:
              for(let t=0;t<pattern_stright_7.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_stright_7[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
            case 7:
              for(let t=0;t<pattern_stright_8.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_stright_8[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
            case 8:
              for(let t=0;t<pattern_stright_9.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_stright_9[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
            case 9:
              for(let t=0;t<pattern_stright_10.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_stright_10[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
            case 10:
              for(let t=0;t<pattern_4_1.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_4_1[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
            case 11:
              for(let t=0;t<pattern_4_2.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_4_2[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
            case 12:
              for(let t=0;t<pattern_4_3.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_4_3[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
            case 13:
              for(let t=0;t<pattern_4_4.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_4_4[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
            case 14:
              for(let t=0;t<pattern_4_5.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_4_5[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
            case 15:
              for(let t=0;t<pattern_4_6.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_4_6[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
            case 16:
              for(let t=0;t<pattern_4_7.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_4_7[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
            case 17:
              for(let t=0;t<pattern_4_8.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_4_8[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
            case 18:
              for(let t=0;t<pattern_4_9.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_4_9[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
            case 19:
              for(let t=0;t<pattern_4_10.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_4_10[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
            case 20:
              for(let t=0;t<pattern_4_11.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_4_11[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
          }
          let total_card = 0;
          total_card = parseInt(data_win.length)
          if(total_card < 5){
            for(let i=0;i<data_array.length;i++){
              if(data_array[i].code_card == "JK"){
                data_win.push(data_array[i])
              }
            }
          }
        
          total_card = parseInt(data_win.length)
          if(total_card == 5){
            flag_func = true;
          }
          
          break;
        }
      }
      if(flag_func == false){
        data_win = [];
      }
      return [flag_func,data_win,6];
    }
    function threeofkind_factory(data_array){
      let flag_func = false
      let data_win = []
      let counts = []
      for(let i=0;i<data_array.length;i++){
        if(counts[data_array[i].val_display]){
          counts[data_array[i].val_display] += 1
        }else{
          counts[data_array[i].val_display] = 1
        }
      }
      let temp = [];
      for(let prop in counts){
        if (counts[prop] >= 2){
              temp.push(prop + ":" + counts[prop])
          }
      }
      if(temp.length > 0){
        let temp_string = temp[0]
        let temp_result = temp_string.split(":");
        let total_temp = temp_result[1];
        let total_jk = 0;
        let total_card = 0;
        
        if(parseInt(total_temp) == 3){
          for(let i=0;i<data_array.length;i++){
            if(data_array[i].code_card == "JK"){
              total_jk = total_jk + 1;
            }
          }
          total_card = parseInt(total_temp) + total_jk
        
          if(total_card == 4){
            info_result = "3 Of A Kind"
            info_card = pattern_stright_10
  
            for(let i=0;i<temp.length;i++){
              temp_string = temp[i]
              temp_result = temp_string.split(":");
              for(let i=0;i<data_array.length;i++){
                if(data_array[i].val_display == temp_result[0]){
                  data_win.push(data_array[i])
                }
              }
            }
  
            // credit_animation(credit,7,totalbet)
            flag_func = true;
          }
          if(total_card == 3){
            info_result = "3 Of A Kind"
            info_card = pattern_stright_10
  
            for(let i=0;i<temp.length;i++){
              temp_string = temp[i]
              temp_result = temp_string.split(":");
              for(let i=0;i<data_array.length;i++){
                if(data_array[i].val_display == temp_result[0]){
                  data_win.push(data_array[i])
                }
              }
            }
  
            // credit_animation(credit,7,totalbet)
            flag_func = true;
          }
        }
        if(parseInt(total_temp) == 2){
          for(let i=0;i<data_array.length;i++){
            if(data_array[i].code_card == "JK"){
              total_jk = total_jk + 1;
            }
          }
          total_card = parseInt(total_temp) + total_jk
          if(total_card == 3){
            info_result = "3 Of A Kind"
            info_card = pattern_stright_10
  
            for(let i=0;i<temp.length;i++){
              temp_string = temp[i]
              temp_result = temp_string.split(":");
              for(let i=0;i<data_array.length;i++){
                if(data_array[i].val_display == temp_result[0]){
                  data_win.push(data_array[i])
                }
                if(data_array[i].val_display == "1"){
                  data_win.push(data_array[i])
                }
              }
            }
            // credit_animation(credit,7,totalbet)
            flag_func = true;
          }
        }
      }
      if(flag_func == false){
        data_win = [];
      }
      return [flag_func,data_win,7]
    }
    function twopair_factory(data_array){
      let flag_func =false
      let data_win = []
      let counts = []
      for(let i=0;i<data_array.length;i++){
        if(counts[data_array[i].val]){
          counts[data_array[i].val] += 1
        }else{
          counts[data_array[i].val] = 1
        }
      }
     
      let temp = [];
      for(let prop in counts){
        if (counts[prop] >= 2){
              temp.push(prop + ":" + counts[prop])
          }
      }
      let total = 0;
      let total_temp = temp.length
      let temp_string = ""
      let temp_result;
      for(let i=0;i<total_temp;i++){
          temp_string = temp[i]
          temp_result = temp_string.split(":");    
          total = total + parseInt(temp_result[1])
      }
      let flag_two = false
      if(total_temp == 3){
        if(total == 4 || total == 6){
          for(let x=0;x<temp.length;x++){
            temp_result = temp[x].split(":");
            switch(temp_result[0]){
              case "10":
                flag_two = true;break;
              case "J":
                flag_two = true;break;
              case "Q":
                flag_two = true;break;
              case "K":
                flag_two = true;break;
              case "AS":
                flag_two = true;break;
              case "JK":
                flag_two = true;break;
            }
          }
          
          if(flag_two){//2 PAIR
            info_result = "2 PAIR"
            info_card = temp
            flag_func = true
  
            for(let i=0;i<temp.length;i++){
              temp_string = temp[i]
              temp_result = temp_string.split(":");
              for(let i=0;i<data_array.length;i++){
                if(data_array[i].val == temp_result[0]){
                  data_win.push(data_array[i])
                }
              }
            }
          }
        }
      }
      if(total_temp == 2){
        if(total == 4 || total == 6){
          for(let x=0;x<temp.length;x++){
            temp_result = temp[x].split(":");
            switch(temp_result[0]){
              case "10":
                flag_two = true;break;
              case "J":
                flag_two = true;break;
              case "Q":
                flag_two = true;break;
              case "K":
                flag_two = true;break;
              case "AS":
                flag_two = true;break;
              case "JK":
                flag_two = true;break;
            }
          }
          
          if(flag_two){//2 PAIR
            info_result = "2 PAIR"
            info_card = temp
            flag_func = true
  
            for(let i=0;i<temp.length;i++){
              temp_string = temp[i]
              temp_result = temp_string.split(":");
              for(let i=0;i<data_array.length;i++){
                if(data_array[i].val == temp_result[0]){
                  data_win.push(data_array[i])
                }
              }
            }
          }
        }
      }
      if(flag_func == false){
        data_win = [];
      }
      return [flag_func,data_win,8]
    }
    function acepair_factory(data_array){
      let flag_func = false
      let data_win = []
      let counts = []
      for(let i=0;i<data_array.length;i++){
        if(counts[data_array[i].val_display]){
          counts[data_array[i].val_display] += 1
        }else{
          counts[data_array[i].val_display] = 1
        }
      }
      
      let temp = [];
      for(let prop in counts){
        if (counts[prop] >= 2){
              temp.push(prop + ":" + counts[prop])
          }
      }
      let total_as = 0;
      let total_jk = 0;
      let total_card = 0;
      if(temp.length > 0){
        let temp_string = temp[0]
        let temp_result = temp_string.split(":");
        let total_temp = temp_result[1];
        for(let i=0;i<data_array.length;i++){
            if(data_array[i].val_display == temp_result[0]){
              data_win.push(data_array[i])
            }
        }
        for(let i=0;i<data_win.length;i++){
            if(data_win[i].val == "AS"){
              total_as = total_as + 1;
            }
            if(data_win[i].val == "JK"){
              total_jk = total_jk + 1;
            }
        }
        
        if(total_as == 2){ // 2 AS
            info_result = "ACE PAIR"
            info_card = temp
            // credit_animation(credit,9,totalbet)
            flag_func = true;
        }
        if(total_jk == 2){ // 2 JK
            let popped = data_win.pop();
            total_jk = data_win.length
            total_as = 0
            total_card = 0
            console.log(total_jk)
            for(let i=0;i<data_array.length;i++){
              if(data_array[i].val == "AS"){
                total_as = total_as + 1
                data_win.push(data_array[i])
              }
            }
            total_card = total_as + total_jk
            if(total_card == 2){ // 1AS + 1JK
              info_result = "ACE PAIR"
              info_card = temp
              flag_func = true;
            }
        }
      }else{
        for(let i=0;i<data_array.length;i++){
            if(data_array[i].val == "AS"){
              total_as = total_as + 1;
              data_win.push(data_array[i])
            }
            if(data_array[i].code_card == "JK"){
              total_jk = total_jk + 1;
              data_win.push(data_array[i])
            }
        }
        total_card = total_as + total_jk
    
        if(total_card == 2){ // 1 as + 1 jk
            info_result = "ACE PAIR"
            info_card = temp
            flag_func = true;
        }
      }
  
      if(flag_func == false){
        data_win = [];
      }
  
      return [flag_func,data_win,9]
    }
    function checkArray(arr_1,arr_2){
      return arr_1.every((val) => arr_2.includes(val))
    }
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    
    
    async function handleSave() {
        let flag = true
        let msg = ""
     
        if(sData == "New"){
            if(pattern_list.length < 1){
                flag = false
                msg += "The ID is required\n"
            }
            
        }else{
            if(idrecord == ""){
                flag = false
                msg += "The ID is required\n"
            }
           
        }
        
        if(flag){
            console.log(pattern_list)
            flag_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/patternsave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page:"CURR-SAVE",
                    pattern_search: searchHome,
                    pattern_page: parseInt(pagingnow),
                    pattern_id: idrecord,
                    pattern_codepoin: resultcodepoint_after ,
                    pattern_status: resultstatus,
                    pattern_resultcardwin: resultcardwin,
                    data: pattern_list,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                flag_btnsave = true;
                if(sData=="New"){
                    // clearField()
                }
                msgloader = json.message;
                RefreshHalaman()
            } else if(json.status == 403){
                flag_btnsave = true;
                alert(json.message)
            } else {
                flag_btnsave = true;
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        }else{
            alert(msg)
        }
    }
    async function handleSaveManual() {
        let flag = true
        let msg = ""
     
        if(pattern_field == ""){
            flag = false
            msg += "The ID is required\n"
        }
       
        if(flag){
            flag_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/patternsavemanual", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: "New",
                    page:"CURR-SAVE",
                    pattern_search: searchHome,
                    pattern_page: parseInt(pagingnow),
                    pattern_id: pattern_field,
                    pattern_idcard: pattern_img,
                    pattern_codepoin: pattern_codepoin ,
                    pattern_resultcardwin: pattern_win,
                    pattern_status: pattern_status_dua,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                flag_btnsave = true;
                if(sData=="New"){
                    // clearField()
                }
                msgloader = json.message;
                RefreshHalaman()
            } else if(json.status == 403){
                flag_btnsave = true;
                alert(json.message)
            } else {
                flag_btnsave = true;
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        }else{
            alert(msg)
        }
    }
    function clearField(){
      idrecord = "";
      resultcard = "";
      resultnmpoint = "";
      resultcardwin = "";
    }
    function callFunction(event){
        switch(event.detail){
            case "NEW":
                NewData("New","","","");
                break;
            case "NEWMANUAL":
                NewManualData("New","","","");
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
                filterHome = [];
                listHome = [];
                const searchdata = {
                  searchHome,
                };
                dispatch("handleSearch", searchdata);
        }  
    };
    function status(e){
        let result = "LOSE"
        if(e == "Y"){
            result = "WIN"
        }
        return result
    }
    function card_img(e,y){
      // console.log(e)      
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
    function card_img_2(e,y){
      if(y == "Y"){
        if(e != ""){
          let data = e.split(",");
          let total_data = e.split(",").length;
          let img_data = "";
          for(let i=0;i<total_data;i++){
            const searchIndex = card_result_data.findIndex((car) => car.id==data[i]);
            // console.log(searchIndex)
            img_data +="<img width='65px' src='"+card_result_data[searchIndex].img+"' /> "
          }
          return img_data
        }else{
          return ""
        }
      }else{
        return ""
      }
    }
    const handleSelectPaging = (event) => {
      let page = event.target.value;
      pagingnow = page;
      const pattern = {
        page,status_search
      };
      dispatch("handlePaging", pattern);
    };
    const handleSelectPagingByPoint = (event) => {
      let page = event.target.value;
      selectPagingPatternByCode = page
      pagePatternByCode = page
      handleCallpatternlist(poin_code)
    };
    const handleSelectSearchStatus = (event) => {
      let searchstatus = event.target.value;
      const pattern = {
        searchstatus,
      };
      dispatch("handleSearchStatus", pattern);
    };
    const call_pattern_list = (e) => {
      pagePatternByCode = 0
      selectPagingPatternByCode = "0"
      listPagePatternByCode = [];
      poin_code = e
      const searchIndex = list_point.findIndex((card) => card.code==e);
      poin_name = list_point[searchIndex].name
      handleCallpatternlist(e)
      myModal_newentry = new bootstrap.Modal(document.getElementById("modalpatternbycode"));
      myModal_newentry.show();
    };
    async function handleCallpatternlist(e) {
        listPatternByCode = [];
        
        const res = await fetch("/api/patterncode", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                pattern_code: e,
                pattern_page : parseInt(pagePatternByCode)
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            perpagePatternByCode = json.perpage;
            totalrecordallPatternByCode = json.totalrecord;
            if (record != null) {
                let totalpaging = Math.ceil(parseInt(totalrecordallPatternByCode) / parseInt(perpagePatternByCode))
                let no = 0
                if(pagePatternByCode > 1){
                    no = parseInt(pagePatternByCode) 
                }
                for (var i = 0; i < record.length; i++) {
                    no = parseInt(no) + 1;
                    listPatternByCode = [
                        ...listPatternByCode,
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
                listPagePatternByCode = [];
                for(var i=1;i<=totalpaging;i++){
                  listPagePatternByCode = [
                        ...listPagePatternByCode,
                        {
                            page_id: i,
                            page_value: ((i*perpagePatternByCode)-perpagePatternByCode),
                            page_display: i + " Of " + perpagePatternByCode*i,
                        },
                    ];
                }
            }
            
        } else {
            logout();
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
                button_function="NEW"
                button_title="New"
                button_css="btn-dark"/>
            <Button
                on:click={callFunction}
                button_function="NEWMANUAL"
                button_title="New Manual"
                button_css="btn-dark"/>
            <Button
                on:click={callFunction}
                button_function="REFRESH"
                button_title="Refresh"
                button_css="btn-primary"/>
            <Panel
                card_search={true}
                card_title="{title_page}"
                card_footer={totalrecord}>
                <slot:template slot="card-title">
                  <div class="float-end">
                    <select
                      on:change={handleSelectPaging}
                      style="text-align: center;"
                      class="form-control">
                      {#each listPage as rec}
                        <option value={rec.page_value}>{rec.page_display}</option>
                      {/each}
                    </select>
                  </div>
                </slot:template>
                <slot:template slot="card-search">
                    <div class="col-lg-12" style="padding: 5px;">
                      <div class="alert alert-primary" role="alert">
                        <table class="table">
                          <tr>
                            <td>
                              <table style="width: 100%;">
                                <tr>
                                  <td>Total Win</td>
                                  <td style="text-align: right;">{totalwin}</td>
                                </tr>
                                <tr>
                                  <td>Total Lose</td>
                                  <td style="text-align: right;">{totallose}</td>
                                </tr>
                                <tr>
                                  <td>Total Pattern</td>
                                  <td style="text-align: right;">{parseInt(totalwin)+parseInt(totallose)}</td>
                                </tr>
                                <tr>
                                  <td style="cursor: pointer;text-decoration: underline;" on:click={() => {
                                      call_pattern_list("RF");
                                  }}>Royal Flush</td>
                                  <td style="text-align: right;">{point_royalflush}</td>
                                </tr>
                                <tr>
                                  <td style="cursor: pointer;text-decoration: underline;" on:click={() => {
                                    call_pattern_list("5K");
                                  }}>5 Of A Kind</td>
                                  <td style="text-align: right;">{point_5ofkind}</td>
                                </tr>
                                <tr>
                                  <td style="cursor: pointer;text-decoration: underline;" on:click={() => {
                                    call_pattern_list("SF");
                                  }}>Straight Flush</td>
                                  <td style="text-align: right;">{point_straightflush}</td>
                                </tr>
                                <tr>
                                  <td style="cursor: pointer;text-decoration: underline;" on:click={() => {
                                    call_pattern_list("4K");
                                    }}>4 Of A Kind</td>
                                  <td style="text-align: right;">{point_4ofkind}</td>
                                </tr>
                                <tr>
                                  <td style="cursor: pointer;text-decoration: underline;" on:click={() => {
                                    call_pattern_list("FH");
                                    }}>Full House</td>
                                  <td style="text-align: right;">{point_fullhouse}</td>
                                </tr>
                              </table>
                            </td>
                            <td>&nbsp;</td>
                            <td>
                              <table style="width: 100%;">
                                <tr><td colspan="2">&nbsp;</td></tr>
                                <tr><td colspan="2">&nbsp;</td></tr>
                                <tr><td colspan="2">&nbsp;</td></tr>
                                <tr>
                                  <td style="cursor: pointer;text-decoration: underline;" on:click={() => {
                                    call_pattern_list("FL");
                                   }}>Flush</td>
                                  <td style="text-align: right;">{point_flush}</td>
                                </tr>
                                <tr>
                                  <td style="cursor: pointer;text-decoration: underline;" on:click={() => {
                                    call_pattern_list("ST");
                                    }}>Straight</td>
                                  <td style="text-align: right;">{point_straight}</td>
                                </tr>
                                <tr>
                                  <td style="cursor: pointer;text-decoration: underline;" on:click={() => {
                                    call_pattern_list("3K");
                                    }}>3 Of A Kind</td>
                                  <td style="text-align: right;">{point_3ofkind}</td>
                                </tr>
                                <tr>
                                  <td style="cursor: pointer;text-decoration: underline;" on:click={() => {
                                    call_pattern_list("2P");
                                    }}>2 Pair (10 Pair)</td>
                                  <td style="text-align: right;">{point_2pair}</td>
                                </tr>
                                <tr>
                                  <td style="cursor: pointer;text-decoration: underline;" on:click={() => {
                                    call_pattern_list("AP");
                                  }}>Ace Pair</td>
                                  <td style="text-align: right;">{point_acepair}</td>
                                </tr>
                              </table>
                            </td>
                          </tr>
                        </table>
                      </div>
                    </div>
                    <div class="row" style="padding: 5px;">
                      <div class="col-lg-8" >
                        <input
                            bind:value={searchHome}
                            on:keypress={handleKeyboard_checkenter}
                            type="text"
                            class="form-control"
                            placeholder="Search Pattern"
                            aria-label="Search"/>
                    </div>
                    <div class="col-lg-4" >
                      <select
                        on:change={handleSelectSearchStatus}
                        bind:value={status_search}
                        class="form-control">
                        <option value="">--Select Status--</option>
                        <option value="Y">WIN</option>
                        <option value="N">LOSE</option>
                      </select>
                    </div>
                    </div>
                </slot:template>
                <slot:template slot="card-body">
                    <table class="table table-striped ">
                        <thead>
                            <tr>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;" >&nbsp;</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">STATUS</th>
                                <th NOWRAP width="10%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">POIN</th>
                                <th NOWRAP width="50%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">PATTERN</th>
                                <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">RESULTCARDWIN</th>
                            </tr>
                        </thead>
                        {#if totalrecord > 0}
                        <tbody>
                            {#each filterHome as rec }
                                <tr>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i on:click={() => {
                                                NewData("Edit",rec.home_id, rec.home_card, rec.home_codepoin, rec.home_nmpoin,rec.home_resultcardwin,
                                                rec.home_create,rec.home_update);
                                            }} class="bi bi-pencil"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_no}</td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">
                                        <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec.home_status_css}">
                                            {status(rec.home_status)}
                                        </span>
                                    </td>
                                    <td  NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">
                                      {rec.home_codepoin}-{rec.home_nmpoin}
                                    </td>
                                    <td  NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                                      {rec.home_id} | <span class="badge bg-info text-dark">{status(check_pattern_status(rec.home_id))}</span><br />
                                      {@html card_img(rec.home_card,75)}
                                    </td>
                                    <td  NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                                      {@html card_img_2(rec.home_resultcardwin,rec.home_status)}
                                    </td>
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
	modal_id="modalentrycrud"
	modal_size="modal-dialog-centered modal-lg"
	modal_title="{title_page+"/"+sData}"
  modal_body_css="height:500px; overflow-y: scroll;"
  modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
    {#if sData == "New"}
      <table>
        {#each pattern_list as rec}
        <tr>
            <td>
              {rec.idpattern} | {rec.point} |{rec.status}<br />
              {@html card_img(rec.idcard,75)}
              {#if rec.status == "Y"}
                <br />
                {@html card_img(rec.resultwin,75)}<br />
              {/if}
            </td>
        </tr>
        {/each}
      </table>  
    {:else}
      {idrecord}<br />
      {@html card_img(resultcard,85)}<br /><br />
      Card Win : {resultcodepoint}-{resultnmpoint} <br />
      {@html card_img(resultcardwin,85)}<br /><br />
      <div class="alert alert-info" role="alert">
        Card Win After : {resultcodepoint_after} - {resultnmpoint_after} - {resultstatus}
      </div>
      <div class="mb-3">
        <div class="alert alert-secondary" style="font-size: 11px; padding:10px;" role="alert">
            Create : {create_field}<br />
            Update : {update_field}
        </div>
    </div>
    {/if}
	</slot:template>
	<slot:template slot="footer">
        {#if flag_btnsave}
        <Button on:click={() => {
                handleSave();
            }} 
            button_function="SAVE"
            button_title="Save"
            button_css="btn-warning"/>
        {/if}
	</slot:template>
</Modal>

<Modal
	modal_id="modalentrycrudmanual"
	modal_size="modal-dialog-centered modal-lg"
	modal_title="PATTERN MANUAL"
  modal_body_css="height:500px; overflow-y: scroll;"
  modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Pattern</label>
            <Input bind:value={pattern_field}
                class="required"
                type="text"
                placeholder="Pattern"/>
        </div>
        {#if pattern_img != ""}
        {@html card_img(pattern_img,85)}<br /><br />

        SORT : ASC<br />
        {@html card_img(pattern_img_clone,85)}<br /><br />
        Card Win : {pattern_status_dua} | {pattern_poin} <br />
        {@html card_img(pattern_win,85)}<br />
        {/if}
	</slot:template>
	<slot:template slot="footer">
        <Button on:click={() => {
              handleCheckWinLose();
          }} 
          button_title="Check"
          button_css="btn-info"/>
        {#if flag_btnsavemanual}
        <Button on:click={() => {
                handleSaveManual();
            }} 
            button_title="Save"
            button_css="btn-warning"/>
        {/if}
	</slot:template>
</Modal>

<Modal
	modal_id="modalpatternbycode"
	modal_size="modal-dialog-centered modal-lg"
	modal_title="PATTERN BY {poin_name}"
  modal_body_css="height:500px; overflow-y: scroll;"
  modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
    <div class="float-end">
      <select
        on:change={handleSelectPagingByPoint}
        on:value={selectPagingPatternByCode}
        style="text-align: center;"
        class="form-control">
        {#each listPagePatternByCode as rec}
          <option value={rec.page_value}>{rec.page_display}</option>
        {/each}
      </select>
    </div>
    <table class="table table-striped ">
      <thead>
          <tr>
              <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
              <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">STATUS</th>
              <th NOWRAP width="10%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">POIN</th>
              <th NOWRAP width="50%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">PATTERN</th>
          </tr>
      </thead>
      <tbody>
        {#each listPatternByCode as rec }
            <tr>
                
                <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_no}</td>
                <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">
                    <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec.home_status_css}">
                        {status(rec.home_status)}
                    </span>
                </td>
                <td  NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">
                  {rec.home_codepoin}-{rec.home_nmpoin}
                </td>
                <td  NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                  {rec.home_id} | <span class="badge bg-info text-dark">{status(check_pattern_status(rec.home_id))}</span><br />
                  {@html card_img(rec.home_card,75)}
                  <br />
                  <b>RESULTCARDWIN</b><br />
                  {@html card_img_2(rec.home_resultcardwin,rec.home_status)}
                </td>
            </tr>
        {/each}
      </tbody>
    </table>
	</slot:template>
	<slot:template slot="footer">
    Record : {totalrecordallPatternByCode}
	</slot:template>
</Modal>