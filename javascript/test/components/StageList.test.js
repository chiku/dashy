// javascript/test/components/StageList.test.js
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2017. All rights reserved
// License::   MIT

const expect = require('chai').expect;

const Stage = require('../../src/components/Stage');
const StageList = require('../../src/components/StageList');

describe('StageList', () => {
  describe('#render', () => {
    const stageOne = {
      name: 'Compile',
      status: 'Passed',
    };
    const stageTwo = {
      name: 'Test',
      status: 'Building',
    };
    const stageList = new StageList().render([stageOne, stageTwo]);

    it('creates a DOM representation', () => {
      expect(stageList[0]).to.equal('div');
    });

    it('has CSS class', () => {
      expect(stageList[1]).to.deep.equal({
        class: 'stage-container',
      });
    });

    it('has Stage children', () => {
      const children = stageList[2];
      expect(children).to.have.length(2);

      expect(children[0][0]).to.equal(Stage);
      expect(children[0][1]).to.equal(stageOne);

      expect(children[1][0]).to.equal(Stage);
      expect(children[1][1]).to.equal(stageTwo);
    });
  });
});
