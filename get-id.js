////////////////////////////////////////////////////////////////
// Read your "Docker ID" message on the Tangle               //
// This is a Java Script to get your Docker ID              //
// This first draft is based on the code of the IOTA guide //
////////////////////////////////////////////////////////////

const Iota = require('@iota/core');
const Extract = require('@iota/extract-json');

// Connect to a node
const iota = Iota.composeAPI({
  provider: 'https://nodes.devnet.thetangle.org:443'
});

// Define the bundle hash whose transactions you want to get
const bundle =
  'SIHQISXRUHFGZBCHOQLRYFXYTQBIERIJZHCHUUJZPAZC9YEQQVXAJFZNZKEBKPILI9GHYX9QCPAYGFWDD';


// Get the transaction objects in the bundle
iota.findTransactionObjects({ bundles: [bundle] })
  .then(bundle => {
    console.log(JSON.parse(Extract.extractJson(bundle)));
  })
  .catch(err => {
    console.error(err);
  });
