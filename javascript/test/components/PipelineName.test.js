// javascript/test/components/PipelineName.test.js
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2017. All rights reserved
// License::   MIT

const expect = require('chai').expect;

const PipelineName = require('../../src/components/PipelineName');

describe('PipelineName', () => {
  describe('#render', () => {
    const stage = new PipelineName().render('Dashy');

    it('creates a DOM representation', () => {
      expect(stage[0]).to.equal('div');
    });

    it('has CSS class', () => {
      expect(stage[1]).to.deep.equal({
        class: 'pipeline-name',
      });
    });

    it('has contents based on its name', () => {
      expect(stage[2]).to.equal('Dashy');
    });
  });
});
