const fs = require('fs');
const readline = require('readline');

function pointsFromChoiceOfTool(tool) {
  switch(tool) {
    case "X":
      return 1
    case "Y":
      return 2
    case "Z":
      return 3
    default:
      console.log("error")
  }
}

function normalise(choice) {
  switch(choice) {
    case "A":
      return "X"
    case "B":
      return "Y"
    case "C":
      return "Z"
  }
}

function normaliseInverseWin(opponent) {
  switch(opponent) {
    case "A":
      return "Y"
    case "B":
      return "Z"
    case "C":
      return "X"
  }
}

function normaliseInverseLoss(opponent) {
  switch(opponent) {
    case "A":
      return "Z"
    case "B":
      return "X"
    case "C":
      return "Y"
  }
}

function pointsFromOutcome(opponent, player) {

    if (opponent == player){
      return 3
    }
    if (opponent=="X" && player=="Y"){
      return 6
    }
    if (opponent=="Y"&&player=="Z"){
      return 6
    }
    if (opponent=="Z" && player=="X"){
      return 6
    }

    return 0
  }

  function pointsFromPartTwoOutcome(opponent, outcome) {
    // z is win, x is lose, y is draw
    if (outcome=="X") {
      return pointsFromChoiceOfTool(normaliseInverseLoss(opponent))
    }

    if (outcome=="Y") {
      return 3 + pointsFromChoiceOfTool(normalise(opponent))
    }

    if (outcome=="Z") {
      return 6 + pointsFromChoiceOfTool(normaliseInverseWin(opponent))
    }

  }


function main() {
  var points = 0
  var pointsPartTwo = 0
  var r = readline.createInterface({
      input : fs.createReadStream("input.txt")
  });
  r.on('line', function (text) {
  points += pointsFromChoiceOfTool(text[2])
  points += pointsFromOutcome(normalise(text[0]), text[2])
  pointsPartTwo += pointsFromPartTwoOutcome(text[0], text[2])
  console.log(points)
  console.log(pointsPartTwo)
  });
}

main()