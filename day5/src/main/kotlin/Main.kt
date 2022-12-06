import java.io.File

fun main() {
    var instructions: MutableList<String> = mutableListOf()
    var initialStackSchema: MutableList<String> = mutableListOf()

    var reachedInstructions = false
    for (l in readFileAsLinesUsingReadLines("./src/main/resources/input.txt")) {
        if (l == "") {
            reachedInstructions = true
            continue
        }

        if (reachedInstructions) {
            instructions.add(l)
            continue
        }

        initialStackSchema.add(l)
    }

    // parse instructions
    val parsedInstructions = parseInstructions(instructions)

    // parse initialStackSchema
    val parsedInitialStackSchema = parseStackSchema(initialStackSchema)
    val parsedInitialStackSchemaForPartTwo = parseStackSchema(initialStackSchema)
    // apply parsed instructions to stacks
    for (i in parsedInstructions) {
        manipulateStacks(i, parsedInitialStackSchema)
        manipulateStacksForPartTwo(i, parsedInitialStackSchemaForPartTwo)
    }
    println("Part One: " + getResult(parsedInitialStackSchema))

    println("Part Two: " + getResult(parsedInitialStackSchemaForPartTwo))
}

fun readFileAsLinesUsingReadLines(fileName: String): List<String> = File(fileName).readLines()

fun getResult(stacks: MutableList<Box>) =
    stacks.joinToString(separator = "") { it.stack[it.stack.lastIndex] }

fun parseInstructions(instructions: List<String>): List<Instruction> {
    val parsedInstructions = mutableListOf<Instruction>()
    for (i in instructions) {
        val moves = findNumbersRegex(i)
        val stackNumberValues = findNumbersRegex(moves).split(",").map {
            it
        }
        val instruction = Instruction(
            stackNumberValues[0].toInt(),
            stackNumberValues[1].toInt() - 1,
            stackNumberValues[2].toInt() - 1
        )
        parsedInstructions.add(instruction)
    }

    return parsedInstructions
}

fun findNumbersRegex(i: String): String {
    return Regex("\\d{1,2}").findAll(i).map { it.groupValues[0] }.joinToString(separator = ",")
}

fun parseStackSchema(initialStackSchema: List<String>): MutableList<Box> {
    val stackNumbers = initialStackSchema[initialStackSchema.lastIndex]
    val stackNumberValues = findNumbersRegex(stackNumbers).split(",").map {
        it
    }
    val stackNumberValueBoxIndexes = stackNumberValues.map { it.toInt().dec() }
    val indexes = mutableListOf<Int>()
    for (int in stackNumberValues) {
        indexes.add(stackNumbers.indexOf(int))
    }
    val map: Map<Int, Int> = indexes.zip(stackNumberValueBoxIndexes).toMap()
    val numberOfStacks =
        stackNumberValues[stackNumberValues.size - 1]

    val stacks = (1..numberOfStacks.toInt()).map { Box(mutableListOf()) }.toMutableList()
    for (i in initialStackSchema) {
        for (j in indexes) {
            if (i[j].isLetter()) {
                stacks[map[j] ?: throw Exception("j not found in map")].stack.add(0, i[j].toString())
            }
        }
    }

    return stacks
}

fun manipulateStacks(instructions: Instruction, stacks: MutableList<Box>): List<Box> {
    repeat(instructions.numberToMove) {
        stacks[instructions.boxIndexTo].stack.add(
            stacks[instructions.boxIndexFrom].stack.removeLast()
        )
    }

    return stacks
}

fun manipulateStacksForPartTwo(instructions: Instruction, stacks: MutableList<Box>): List<Box> {
    when (instructions.numberToMove) {
        1 -> manipulateStacks(instructions, stacks)
        else -> manipulateMultipleStacks(instructions, stacks)
    }
    return stacks
}

fun manipulateMultipleStacks(instructions: Instruction, stacks: MutableList<Box>): List<Box> {
    val stackToAppend: MutableList<String> = mutableListOf()
    repeat(instructions.numberToMove) {
        stackToAppend.add(0, stacks[instructions.boxIndexFrom].stack.removeLast())
    }
    stacks[instructions.boxIndexTo] =
        Box((stacks[instructions.boxIndexTo].stack + stackToAppend) as MutableList<String>)
    return stacks
}

data class Instruction(
    val numberToMove: Int,
    val boxIndexFrom: Int,
    val boxIndexTo: Int
)

data class Box(
    val stack: MutableList<String>
)
