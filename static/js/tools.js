"use strict"

const toolsEl = document.getElementById("tools")
const numbersEl = document.getElementById("numbers")
const digitEls = numbersEl.querySelectorAll("div[data-digit]")
const firstDigitEls = numbersEl.querySelectorAll('[data-digit="1"]')
const secondDigitEls = numbersEl.querySelectorAll('[data-digit="2"]')
const thirdDigitEls = numbersEl.querySelectorAll('[data-digit="3"]')
const numbersDigitEls = document.querySelectorAll("#numbers > div > div")
const roundsEl = document.getElementById("rounds")
const roundsBtnEl = document.getElementById("rounds_btns")
const addRoundBtnEl = document.getElementById("new_round")
const roundTemplateEl = document.getElementById("round_template").content.querySelector("div")
const makeGuessBtnEl = document.getElementById("check_solution_btn")
const checkSolutionsBtnEl = document.getElementById("check_solution_btn")
const solveDialogEl = document.getElementById("solve")
const solutionEl = document.getElementById("solution")
const numberOfRoundsEl = solveDialogEl.querySelector(".number_of_rounds")
const numberOfQuestionsEl = solveDialogEl.querySelector(".number_of_questions")
const checkSolutionEl = document.getElementById("check_solution")
const solutionTemplateEl = document.getElementById("solution_template").content.querySelector("div")

/**
 * Removes all the rounds
 */
function clear_rounds() {
    roundsEl.querySelectorAll("#rounds > div:not(#rounds_btns)").forEach((el)=>{
        el.remove()
    })
}

/**
 * Stores a representation of the tools' state
 */
function store_tools_state() {
    const state = {
        digits: [],
        rounds: [],
    }
    // Get digits state
    digitEls.forEach((el)=>{
        if (el.classList.contains("red")) {
            state.digits.push("red")
        } else if (el.classList.contains("green")) {
            state.digits.push("green")
        } else {
            state.digits.push("blank")
        }
    })
    // Get rounds state
    roundsEl.querySelectorAll("#rounds > div:not(#rounds_btns)").forEach((el)=>{
        if (el.classList.contains("round")) {
            // It's a round
            let digits = []
            el.querySelectorAll("select").forEach((digit)=>{
                digits.push(digit.value)
            })
            let cards = []
            el.querySelectorAll("[data-card]").forEach((card)=>{
                let card_class = false
                if (card.classList.contains("red")) {
                    card_class = "red"
                } else if (card.classList.contains("green")) {
                    card_class = "green"
                }
                cards.push({
                    class: card_class,
                    verified: (card.dataset.verified ? card.dataset.verified : false),
                })
            })
            state.rounds.push({
                type: "round",
                digits: digits,
                cards: cards,
                number_of_cards: el.dataset.cards,
                locked: el.dataset.locked,
            })
        } else {
            // It's a solution
            state.rounds.push({
                type: "solution",
                code: el.querySelector(".solution_code").textContent,
                result: el.querySelector(".solution_result").textContent,
                class: (el.classList.contains("green") ? "green": "orange")
            })
        }
    })
    // Store
    store_object("tools_state", state)
}

/**
 * Checks if there's a tools state
 * @returns true if a tools state is preset
 */
function has_tools_state() {
    return load("tools_state") !== null
}

/**
 * Restores the tools' state if any
 */
function restore_tools_state() {
    const state = load_object("tools_state")
    if (state === null) return
    // Restore digits
    for (let digit = 0; digit < 15; digit++) {
        digitEls[digit].classList.remove("red")
        digitEls[digit].classList.remove("green")
        if (state.digits[digit] === "blank") continue
        digitEls[digit].classList.add(state.digits[digit])
    }
    // Restore rounds
    clear_rounds()
    for (let round of state.rounds) {
        switch (round.type) {
            case "solution":
                const solutionBoxEl = solutionTemplateEl.cloneNode(true)
                solutionBoxEl.querySelector(".solution_code").textContent = round.code
                const resultEl = solutionBoxEl.querySelector(".solution_result")
                resultEl.textContent = round.result
                resultEl.classList.add(round.class)
                roundsBtnEl.insertAdjacentElement("beforebegin", solutionBoxEl)
                break;
        
            case "round":
                const roundEl = roundTemplateEl.cloneNode(true)
                roundEl.dataset.cards = round.number_of_cards
                const roundDigitEls = roundEl.querySelectorAll("select")
                for (let i = 0; i < 3; i++) {
                    if (round.digits[i] === "?") continue
                    roundDigitEls[i].value = round.digits[i]
                }
                const cardsEls = roundEl.querySelectorAll("[data-card]")
                for (let i = 0; i < Math.min(round.cards.length, 6); i++) {
                    if (round.cards[i].verified) cardsEls[i].dataset.verified = round.cards[i].verified
                    if (round.cards[i].class) cardsEls[i].classList.add(round.cards[i].class)
                }
                roundEl.dataset.locked = round.locked
                roundsBtnEl.insertAdjacentElement("beforebegin", roundEl)
                break;
        }
    }
    roundsEl.scrollTop = roundsEl.scrollHeight
}

/**
 * Clears the stored state for tools
 */
function clear_tools_state() {
    clear_storage_item("tools_state")
}

/**
 * Handle numbers open
 */
numbersEl.addEventListener("click", (e)=>{
    if (toolsEl.dataset.view == "numbers") return
    e.preventDefault()
    toolsEl.dataset.view = "numbers"
})

/**
 * Handle rounds open
 */
roundsEl.addEventListener("click", (e)=>{
    if (toolsEl.dataset.view == "rounds") return
    e.preventDefault()
    toolsEl.dataset.view = "rounds"
})

/**
 * Adds a listener to an element to transition between green/red/neutral
 * @param {Node} element 
 */
function handle_red_green(element) {
    element.addEventListener("mousedown", (e)=>{
        if (element.classList.contains("red")) {
            element.classList.remove("red")
            element.classList.add("green")
        } else if (element.classList.contains("green")) {
            element.classList.remove("green")
        } else {
            element.classList.add("red")
        }
        // Update state
        store_tools_state()
    })
}

/**
 * Handle digits select
 */
numbersDigitEls.forEach(handle_red_green)

/**
 * Registers the add round action
 */
addRoundBtnEl.addEventListener("click", (e)=>{
    e.preventDefault()
    const round = roundTemplateEl.cloneNode(true)
    const game = load_game()
    round.dataset.cards = (game !== null ? game.criterias.length : 6)
    round.querySelectorAll("div[data-card]").forEach(handle_red_green)
    roundsBtnEl.insertAdjacentElement("beforebegin", round)
    roundsEl.scrollTop = roundsEl.scrollHeight
    // Update state
    store_tools_state()
})

/**
 * Returns object representing the last round played (if any)
 * @returns Node or null
 */
function get_last_round() {
    const el = roundsBtnEl.previousElementSibling
    if (el !== null && el.classList.contains("round")) {
        return el
    }
    return null
}

/**
 * Counts the number of questions marked in a given round
 * @param {Node|null} round_el the ".round" element we want to check
 * @returns the number of questions verified in the round or 0
 */
function count_questions_in_round(round_el) {
    if (round_el === null) return 0
    return round_el.querySelectorAll('.green, .red, [data-verified]').length
}

/**
 * Retrieves the stats object
 * @returns a stats objects with the # of rounds played and # of questions asked
 */
function get_stats() {
    const stats = {
        rounds: 0,
        questions: 0,
    }
    const roundEls = roundsEl.querySelectorAll("#rounds > div")
    for (let i = 0; i < roundEls.length; i++) {
        if (!roundEls[i].classList.contains("round")) break
        let questions = count_questions_in_round(roundEls[i])
        if (questions === 0) continue
        stats.rounds++
        stats.questions += questions
    }
    return stats
}

/**
 * Returns the last proposal
 * @returns the proposal as a string, or false
 */
function get_last_proposal() {
    const round = get_last_round()
    if (!round) {
        show_alert("Start a round first", 3)
        return false
    }
    const code = `${round.querySelector('select[data-digit="1"]').value}${round.querySelector('select[data-digit="2"]').value}${round.querySelector('select[data-digit="3"]').value}`
    if (code.length !== 3 || code.indexOf("?") !== -1) {
        show_alert("Compose your 3-digits proposal for this round", 6)
        return false
    }
    return code
}

/**
 * Locks the last round so that the proposal cannot be changed
 */
function lock_last_round() {
    const round = get_last_round()
    if (!round) return false
    round.dataset.locked = "1"
    // Update state
    store_tools_state()
}

/**
 * Sets the (last) proposal solution
 * @param {string} card the card number [1,6] that has been verified
 * @param {boolean} success result of the verification
 */
function set_last_proposal_solution(card, success) {
    const round = get_last_round()
    if (!round) return
    const option = round.querySelector(`div[data-card="${card}"]`)
    option.dataset.verified = (success ? "green" : "red")
    // Update state
    store_tools_state()
}

/**
 * Resets the number
 */
function reset_numbers() {
    numbersDigitEls.forEach((digit)=>{
        digit.classList.remove("red")
        digit.classList.remove("green")
    })
}

/**
 * Tries to identify what digit the user has picked
 * @param {[]node} digitEls the possible digits for this number
 * @returns a digit as a string or an empty string
 */
function predict_digit(digitEls) {
    let green = true
    let red = true
    digitEls.forEach((digitEl)=>{
        if (digitEl.classList.contains("red")) return
        if (digitEl.classList.contains("green")) {
            if (green === true) {
                // This is the first green digit
                green = digitEl
            } else {
                // A duplicate!
                green = false
            }
        }
        if (red === true) {
            // This is the first non-red digit
            red = digitEl
        } else {
            // A duplicate!
            red = false
        }
    })
    if (green !== true && green !== false) return green.textContent
    if (red !== true && red !== false) return red.textContent
    return ""
}

/**
 * Tries to identify what code the user thinks is right
 * @returns the code as a string or an empty string
 */
function predict_code() {
    let code = predict_digit(firstDigitEls)
    if (code.length === 0) return ""
    code += predict_digit(secondDigitEls)
    if (code.length === 1) return ""
    code += predict_digit(thirdDigitEls)
    if (code.length === 2) return ""
    return code
}

/**
 * Toggle check solution
 */
checkSolutionsBtnEl.addEventListener("mousedown", (e)=>{
    if (checkSolutionEl.classList.contains("disabled")) return
    e.preventDefault()
    const stats = get_stats()
    numberOfRoundsEl.textContent = stats.rounds + " round" + (stats.rounds === 1 ? "" : "s")
    numberOfQuestionsEl.textContent = stats.questions + " question" + (stats.questions === 1 ? "" : "s")
    solutionEl.value = predict_code()
    solveDialogEl.showModal()
})

/**
 * Register close action for solve dialog
 */
solveDialogEl.querySelector(".close_btn").addEventListener("click", (e)=>{
    e.preventDefault()
    solveDialogEl.close()
})

/**
 * Handle check solution
 */
checkSolutionEl.addEventListener("submit", async (e)=>{
    e.preventDefault()
    show_spinner()
    const code = solutionEl.value
    const solutionBoxEl = solutionTemplateEl.cloneNode(true)
    solutionBoxEl.querySelector(".solution_code").textContent = code;
    const solutionResultEl = solutionBoxEl.querySelector(".solution_result")
    const game_code = load_game_code()
    if (game_code === null) {
        show_alert("Errored while checking game code")
        solveDialogEl.close()
        hide_spinner()
        return
    }
    if (code === game_code) {
        solutionResultEl.textContent = "correct"
        solutionResultEl.classList.add("green")
    } else {
        solutionResultEl.textContent = "wrong"
        solutionResultEl.classList.add("orange")
    }
    roundsBtnEl.insertAdjacentElement("beforebegin", solutionBoxEl)
    solveDialogEl.close()
    // Update state
    store_tools_state()
    hide_spinner()
})