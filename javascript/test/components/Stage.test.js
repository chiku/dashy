// javascript/test/components/Stage.test.js
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2017. All rights reserved
// License::   MIT

const expect = require('chai').expect;

const Stage = require('../../src/components/Stage');

describe('Stage', () => {
  describe('#render', () => {
    const stage = new Stage().render({
      name: 'Test',
      status: 'building',
    });

    it('creates a DOM representation', () => {
      expect(stage[0]).to.equal('div');
    });

    it('has CSS class based on its status', () => {
      expect(stage[1]).to.deep.equal({
        class: 'stage building',
      });
    });

    it('has contents based on its name', () => {
      expect(stage[2]).to.equal('Test');
    });

    describe('when status is not in all lower-case', () => {
      const stageUpcase = new Stage().render({
        name: 'Test',
        status: 'Building',
      });

      it('has a lower-name CSS class name', () => {
        expect(stageUpcase[1]).to.deep.equal({
          class: 'stage building',
        });
      });
    });
  });
});
