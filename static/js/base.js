"use strict"

const supported_languages = {
    "br": ["br", "pt"],
    "cns": ["cns", "cn"],
    "cnt": ["cnt", "cn"],
    "cz": ["cz", "cz"],
    "de": ["de", "de"],
    "en": ["en", "uk"],
    "fr": ["fr", "fr"],
    "gr": ["gr", "gr"],
    "hu": ["hu", "hu"],
    "it": ["it", "it"],
    "jp": ["jp", "jp"],
    "kr": ["kr", "kr"],
    "nl": ["nl", "nl"],
    "pl": ["pl", "pl"],
    "ru": ["ru", "ru"],
    "sp": ["sp", "es"],
    "th": ["th", "th"],
    "ua": ["ua", "ua"],
}
const default_language = "en"

const gameIdEl = document.getElementById("game_id")
const spinnerEl = document.getElementById("spinner")
const alertEl = document.getElementById("alert")
const infoBtnEl = document.getElementById("info_btn")
const infoDialogEl = document.getElementById("info")
const languageSelectEl = document.getElementById("language")

/**
 * Clears the local storage
 */
function clear_storage() {
    localStorage.clear()
}

/**
 * Removes a storage item
 * @param {string} key the item's key
 */
function clear_storage_item(key) {
    localStorage.removeItem(key)
}

/**
 * Stores some data in the local storage
 * @param {string} key identifier for the data to be stored
 * @param {string|null} value data to be stored; if nulls clears removes the key
 * @returns 
 */
function store(key, value=null) {
    if (value === null) {
        localStorage.removeItem(key)
        return
    }
    localStorage.setItem(key, value)
}

/**
 * Retrieves some data from storage
 * @param {string} key identifier for the data to retrieve
 * @returns string|null
 */
function load(key) {
    return localStorage.getItem(key)
}

/**
 * Stores some data in the local storage
 * @param {string} key identifier for the data to be stored
 * @param {any|null} value data to be stored; if nulls clears removes the key
 * @returns 
 */
function store_object(key, value=null) {
    if (value === null) {
        store(key)
        return
    }
    store(key, JSON.stringify(value))
}

/**
 * Retrieves some data from storage
 * @param {string} key identifier for the data to retrieve
 * @returns any|null
 */
function load_object(key) {
    const obj = load(key)
    if (obj === null) return null
    return JSON.parse(obj)
}

/**
 * Checks if data for a given key exists in storage
 * @param {string} key identifier of the data to look for
 * @returns true if key has data, false otherwise
 */
function key_exists(key) {
    return load(key) !== null
}

/**
 * Stores the game object that's currently being played
 * @param {object} game 
 */
function store_game(game) {
    store_object("game", game)
}

/**
 * Returns true if there is a game in storage
 * @returns true if there's a game in storage, false otherwise
 */
function has_game() {
    return load("game") !== null
}

/**
 * Loads the current game
 * @returns the game object or null
 */
function load_game() {
    return load_object("game")
}

/**
 * Fetches the game code
 * @returns the game code as a string or null if no game is stored
 */
function load_game_code() {
    const game = load_object("game")
    if (game === null) return null
    return game.code
}

/**
 * Returns the game id from storage (if any)
 * @param {boolean} pretty if true returns a formatted game id
 * @returns game id as a string or null
 */
function load_game_id(pretty = false) {
    const game = load_object("game")
    if (game === null) return null
    if (pretty) return game.id.replace(/(.{3})/g,"$1 ").trim()
    return game.id
}

/**
 * Retrieves the law for a given card
 * @param {string|number} card the card that we want to retrieve a law for
 * @returns the law or null
 */
function load_law(card) {
    const game = load_object("game")
    if (game === null) return null
    return game.laws[Number(card)-1]
}

/**
 * Stores a checkpoint object
 * @param {any} checkpoint 
 */
function store_checkpoint(checkpoint) {
    store_object("checkpoint", checkpoint)
}

/**
 * Loads a checkpoint object
 * @returns any checkpoint object
 */
function load_checkpoint() {
    return load_object("checkpoint")
}

/**
 * Retrieves the user language
 * @returns a string representing the user language
 */
function get_browser_language_code() {
    const lang_code = navigator.language.substring(0, 2).toLowerCase()
    for (const lang in supported_languages) {
        if (lang == lang_code || supported_languages[lang].indexOf(lang_code) != -1) {
            console.log(`Detected language '${lang}'`)
            return lang
        }
    }
    console.log(`Using default language '${default_language}'`)
    return default_language
}

/**
 * Sets the language code for this player
 * @param {string} lang the tm language code
 */
function set_language_code(lang) {
    store("language", lang)
}

/**
 * Gets the language code to use
 * @returns the language code as a string
 */
function get_language_code() {
    // Fast path
    let lang = load("language")
    if (lang === null) {
        // Slow path
        lang = get_browser_language_code()
        set_language_code(lang)
    }
    return lang
}

/**
 * Add event listener for copy hash (required https)
 */
gameIdEl.addEventListener("click", (e)=>{
    e.preventDefault()
    navigator.clipboard.writeText(window.location)
})

/**
 * Returns the game id from the url hash
 * @param {boolean} pretty if true returns a formatted game id
 * @returns game id as a string (or an empty string)
 */
function get_hash_id(pretty = false) {
    if (!pretty) return window.location.hash.replace("#", "").trim()
    return window.location.hash.replace("#", "").replace(/(.{3})/g,"$1 ").trim()
}

/**
 * Sets the game ID hash in url hash and header
 */
function set_game_id() {
    const game_id = load_game_id(false)
    const game_id_pretty = load_game_id(true)
    window.location.hash = (game_id ? game_id : "")
    gameIdEl.textContent = (game_id_pretty ? game_id_pretty : "?")
}

/**
 * Shows the spinner modal
 */
function show_spinner() {
    spinnerEl.showModal()
}

/**
 * Hides the spinner modal
 */
function hide_spinner() {
    spinnerEl.close()
}

let alert_timeout;

/**
 * Shows the alert if no other alerts are present
 * @param {string} msg the message to display
 * @param {number} timeout in seconds
 * @returns 
 */
function show_alert(msg, timeout=3) {
    clearTimeout(alert_timeout)
    alertEl.textContent = msg
    alertEl.open = true
    alert_timeout = setTimeout(()=>{
        alertEl.close()
    }, timeout * 1000)
}

/**
 * Allow the alert to be closed manually
 */
alertEl.addEventListener("click", (e)=>{
    e.preventDefault()
    if (alert_timeout) clearTimeout(alert_timeout)
    alertEl.close()
})

/**
 * Shows the info modal
 */
function show_info() {
    const game = load_game()
    const game_id = load_game_id(true)
    let criterias = "?"
    let verifiers = "?"
    if (game !== null) {
        criterias = game.criterias.join(", ")
        verifiers = game.verifiers.join(", ")
    }
    languageSelectEl.value = get_language_code()
    infoDialogEl.querySelector(".game_id").textContent = (game_id ? game_id : "?")
    infoDialogEl.querySelector(".game_criterias").textContent = criterias
    infoDialogEl.querySelector(".game_vc").textContent = verifiers
    infoDialogEl.showModal()
}

/**
 * Register the handler for the info button
 */
infoBtnEl.addEventListener("click", (e)=>{
    e.preventDefault()
    show_info()
})

/**
 * Register language change
 */
languageSelectEl.addEventListener("change", async (e)=>{
    e.preventDefault()
    set_language_code(languageSelectEl.value)
    show_spinner()
    await update_cards_language()
    hide_spinner()
})

/**
 * Adds a handler to close the info dialog
 */
infoDialogEl.querySelector(".close_btn").addEventListener("click", (e)=>{
    e.preventDefault()
    infoDialogEl.close()
})