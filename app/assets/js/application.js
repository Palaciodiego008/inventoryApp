require("bootstrap/dist/js/bootstrap.bundle.js");
require("bootstrap/dist/js/bootstrap.bundle.min.js");
require("@fortawesome/fontawesome-free/js/all.js");

import { Application } from "@hotwired/stimulus"
import { definitionsFromContext } from "@hotwired/stimulus-webpack-helpers"

const app = Application.start()
const context = require.context("./controllers", true, /\.js$/)
app.load(definitionsFromContext(context))
