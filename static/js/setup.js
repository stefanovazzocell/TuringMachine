"use strict"

const welcomeDialogEl = document.getElementById("welcome")

const continuePlayingEl = document.getElementById("continue_playing")
const continueEl = document.getElementById("continue")
const restoredGameIdEl = document.getElementById("restored_game_id")

const startGameEl = document.getElementById("start_game")
const gameIdInputEl = document.getElementById("game_id_input")

const createGameEl = document.getElementById("create_game")
const createGameDifficultyEl = document.getElementById("create_game_difficulty")
const createGameCardsEl = document.getElementById("create_game_cards")

const searchGameEl = document.getElementById("search_game")
const criteriasInputEl = document.getElementById("criterias_input")
const verifiersInputEl = document.getElementById("verifiers_input")

const newGameBtnEl = document.getElementById("new_game")

const duplicateDialogEl = document.getElementById("duplicate_game")

/**
 * Show welcome screen
 */
function show_welcome() {
    const game_id = (get_hash_id(true) || load_game_id(true))
    restoredGameIdEl.textContent = game_id
    gameIdInputEl.value = game_id
    if (has_game()) {
        continuePlayingEl.classList.remove("disabled")
    } else continuePlayingEl.classList.add("disabled")
    welcomeDialogEl.showModal()
}

/**
 * Closes the welcome screen
 */
function close_welcome() {
    welcomeDialogEl.close()
}

/**
 * Starts a game
 * @param {object} game the game object returned by the API
 */
async function start_game(game) {
    store_game(game)

    const cards = set_cards(game.criterias, game.verifiers)
    clear_rounds()
    reset_numbers()
    set_game_id()

    // Wait for cards to be loaded
    await cards
}

/**
 * Starts a game, clears any stored state, closes the welcome screen and hides
 * the spinner.
 * @param {object} game the game object returned by the API
 */
async function new_game(game) {
    const game_loader = start_game(game)
    clear_cards_state()
    clear_tools_state()
    await game_loader
    close_welcome()
    hide_spinner()
}

/**
 * Adds event listener for "continue game"
 */
continueEl.addEventListener("submit", async (e)=>{
    e.preventDefault()
    show_spinner()

    const game_loader = start_game(load_game())
    restore_cards_state()
    restore_tools_state()
    await game_loader

    close_welcome()
    hide_spinner()
})

/**
 * Handler for start game
 */
startGameEl.addEventListener("submit", async (e)=>{
    e.preventDefault()
    show_spinner()
    const game_id = gameIdInputEl.value.replaceAll(" ", "").replaceAll("-", "")
    // Fast path
    if (load_game_id() === game_id) {
        await new_game(load_game())
        return
    }
    // Slow path
    const request = fetch_game(game_id)
    const response = await request
    if (response === false) {
        show_alert("Invalid game ID")
        hide_spinner()
        return
    }
    await new_game(response)
})

/**
 * Handler for create game
 */
createGameEl.addEventListener("submit", async (e)=>{
    e.preventDefault()
    const request = fetch_game(false, createGameDifficultyEl.value, Number(createGameCardsEl.value))
    show_spinner()
    const response = await request
    if (response === false) {
        show_alert("Something went wrong, try again later :(")
        hide_spinner()
        return
    }
    await new_game(response)
})

/**
 * Handler for search game
 */
searchGameEl.addEventListener("submit", async (e)=>{
    e.preventDefault()
    let criterias = criteriasInputEl.value.replaceAll(/,\s+/g, ",").split(",").map((el)=>{ return Number(el) })
    let verifiers = verifiersInputEl.value.replaceAll(/,\s+/g, ",").split(",").map((el)=>{ return Number(el) })
    if (criterias.length != verifiers.length) {
        show_alert("There should be a verifier for each criteria")
        return
    }
    const request = solve_game(criterias, verifiers)
    show_spinner()
    const response = await request
    if (response === false) {
        show_alert("Something went wrong, try again later :(")
        hide_spinner()
        return
    }
    if (response.solutions.length === 0) {
        show_alert("This game has no solution")
        hide_spinner()
        return
    }
    if (response.solutions.length !== 1) {
        show_alert("This game has more than one possible solution")
        hide_spinner()
        return
    }
    response.code = response.solutions[0]
    await new_game(response)
})

/**
 * Registers the action for the "new game" button
 */
newGameBtnEl.addEventListener("click", (e)=>{
    e.preventDefault()
    show_welcome()
})

/**
 * Detect if there is another instance of the game was opened
 */
let tab_code = Math.random().toString().substring(2)
store("tab", tab_code)
const duplicate_game_monitor = setInterval(()=>{
    if (load("tab") === tab_code) return
    // Close all active modals
    document.querySelectorAll(":modal").forEach((el)=>{
        el.close()
    })
    // Show notice
    duplicateDialogEl.showModal()
    // Stop checking
    clearInterval(duplicate_game_monitor)
}, 1000)

/**
 * Main function called when loading is complete
 */
show_welcome()