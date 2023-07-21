const screeps2go = require('screeps2go');
const bytecode = require('wasmbin');

screeps2go.init(bytecode);

module.exports.loop = function() {
  screeps2go.loop();
}
