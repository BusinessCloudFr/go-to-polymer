var user;
var countryList = [];
var matchList = [];
var ranking =[];
var winners = []
var nbBets = 8;
var userService;
var countryService;
var matchService;

$(document).ready(function(){
  $(".check").hide();
  $(".userRanking").hide();

    //display session
    $(".username").text($.session.get("session"));
    
    userService = $("user-service").get(0);
    countryService = $("country-service").get(0);
    matchService = $("match-service").get(0);
    //load all apis
    loadApis().done(function(){
      //get countries and user, store them in global variables
      $.when(getCountries(), getUserByPseudo()).done(function(resultCountries, resultUser){
        user = resultUser;
        countryList = resultCountries;
        createMatches();
        rankingDatastore();
      });
    });      

    //submit bets
    $(".check").click(function(){
      nbBets /= 2;
      createNewMatches(winners);      
    });

    //check if all matches are betted on, if yes show check button
    $(".betsContainer").click(function(event){
      var button = $(event.target);
      if(button.is("paper-material")){
        winners = [];
        $("bet-box").each(function(){
          var result = $(this).get(0).winner;
          if(result != undefined){
            winners.push(result);  
          }        
        });
        if(winners.length == nbBets){
          $(".check").show();
        }
      }
    });
});


  //functions

  // load all apis
  function loadApis(){
    var deferred = $.Deferred();
    $.when(userService.load(), countryService.load(), matchService.load()).done(function(){
      console.log("all apis loaded");
      deferred.resolve();    
    });
    return deferred.promise();
  }

  // get all countries form the datastore
  function getCountries(){
    var deferred = $.Deferred();
    countryService.list().done(function(result){
      deferred.resolve(result);
    });
    return deferred.promise();
  }

  // create initial sets of matches
  function createMatches(){
    $(".load1").hide();
    for(i = 0; i<countryList.countries.length; i=i+2){
      matchList.push({"uidCountryA" : countryList.countries[i].uid, "uidCountryB" : countryList.countries[i+1].uid, "uidUser" : user.uid, "round": 1});
      var betBox = document.createElement("bet-box");
      betBox.countryA = countryList.countries[i];
      betBox.countryB = countryList.countries[i+1];
      $(".betsContainer").append(betBox);
    }
  }

  // create new sets of matches with the winners of the previous round
  function createNewMatches(winners){
    var round = matchList[matchList.length-1].round;
    $(".check").hide();
    setWinnersUid(winners);
    if(round == 4){
      $(".betsContainer").empty();
      submitRanking().done(function(){
        rankingLocal();
      });
    }
    else{
      $(".betsContainer").empty();
      for(i = 0; i<winners.length; i=i+2){
        matchList.push({"uidCountryA" : winners[i].uid, "uidCountryB" : winners[i+1].uid, "uidUser" : user.uid, "round": round+1});
        var betBox = document.createElement("bet-box");
        betBox.countryA = winners[i];
        betBox.countryB = winners[i+1];
        $(".betsContainer").append(betBox);
      }
    }    
  }

  // set winner uid in matches
  function setWinnersUid(winners){
    var i;
    for(i = 0; i<winners.length; i++){
      var j;
      for(j = matchList.length-1; j>=0; j--){
        if(matchList[j].uidCountryA == winners[i].uid || matchList[j].uidCountryB == winners[i].uid){
          matchList[j].uidWinner = winners[i].uid;
        }
      }
    }
    
  }

  // display user ranking with local matches list 
  function rankingLocal(){
    var rankings = {};
    rankings.matchs = matchList;
    console.log(rankings);
    ranking(rankings);
  }

  // display user ranking with datastore
  function rankingDatastore(){
    getRanking().done(function(rankingResult){
      ranking(rankingResult);
    }); 
  }

  // display user ranking
  function ranking(ranking){
    $(".load2").show();
    $(".userRanking").hide();
    // delete previous ranking
    $(".ranksUser").each(function(){
      $(this).children("div").empty();
    });
    
    var length = ranking.matchs.length;
    var i; 
    for(i = 0; i<ranking.matchs.length; i++){
      match = ranking.matchs[i];
      // get only the matches of the current user
      if(user.uid == match.uidUser){

        var countryBox = document.createElement("country-box");
        var countryA = getCountryWithUid(match.uidCountryA);
        var countryB = getCountryWithUid(match.uidCountryB);
        countryBox.countryA = countryA;
        countryBox.countryB = countryB;

        // append matches to the proper category depending of the round
        if(match.round == 4){
          $(".final").children("div").append(countryBox);

          var winner = getCountryWithUid(match.uidWinner)
          var winnerFlag = $("<div class='country'><iron-image style='width:200px; height:150px;' sizing='cover' src=img/"+winner.urlflag+"></iron-image></div>");
          var winnerName = $("<p>"+winner.label+"</p>");
          $(".winner").children("div").append(winnerFlag);
          $(".country").append(winnerName);
        }
        else if(match.round == 1){
          $(".roundOf16").children("div").append(countryBox); 
        }
        else if(match.round == 2){
          $(".quarterFinals").children("div").append(countryBox); 
        }
        else if(match.round == 3){
          $(".semiFinals").children("div").append(countryBox); 
        }
      }
    }  
    $(".load2").hide();
    $(".userRanking").show();
  }

  // save the user ranking list to the datastore
  function submitRanking(){
    var deferred = $.Deferred();
    console.log(matchList);
    matchService.submitRanking(matchList).done(function(result){
      console.log("OK");
      deferred.resolve(result);
    });
    return deferred.promise();
  }

  // get all the rankings from the datastore
  function getRanking(){
    var deferred = $.Deferred();
    matchService.list().done(function(result){
      console.log(result);
      deferred.resolve(result);
    });
    return deferred.promise();
  }

  // get a country object from the datastore with it's uid
  function getCountryWithUid(uid){
    var i;
    for(i = 0; i<countryList.countries.length;i++){
      if(countryList.countries[i].uid == uid){
        return countryList.countries[i];
      }
    }
  }

  // get user object from the datastore with it's username
  function getUserByPseudo(){
    var currentUser = $.session.get("session");
    var deferred = $.Deferred();
    if(currentUser != undefined){
      userService.getUserByPseudo({"pseudo" : currentUser}).done(function(res){
        if(res == null){
          deferred.resolve({uid : ""})
        }
        else{
          deferred.resolve(res);  
        }        
      });
    }
    else{
      deferred.resolve({uid : ""});
    }
    return deferred.promise();
  }