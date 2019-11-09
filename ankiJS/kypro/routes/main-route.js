const express = require('express');

const mainController = require('../controllers/main-controller');

const router = express.Router();

router.get('/', mainController.getFetchFile);

router.post('/post-file', mainController.postFetchFile);

router.get('/bonds', mainController.getBondsWorking);

router.get('/show-value/:position', mainController.getShowValue);

module.exports = router;