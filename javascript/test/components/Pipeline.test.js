// javascript/test/components/Pipeline.test.js
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2017. All rights reserved
// License::   MIT

const expect = require('chai').expect;

const PipelineName = require('../../src/components/PipelineName');
const StageList = require('../../src/components/StageList');
const Pipeline = require('../../src/components/Pipeline');

describe('Pipeline', () => {
  describe('#render', () => {
    const stageOne = {
      name: 'Compile',
      status: 'Passed',
    };
    const stageTwo = {
      name: 'Test',
      status: 'Building',
    };
    const pipeline = new Pipeline().render({
      name: 'Dashy',
      stages: [stageOne, stageTwo],
    });

    it('creates a DOM representation', () => {
      expect(pipeline[0]).to.equal('div');
    });

    it('has CSS class', () => {
      expect(pipeline[1]).to.deep.equal({
        class: 'pipeline',
      });
    });

    it('has DOM children', () => {
      const children = pipeline[2];
      expect(children).to.have.length(2);
    });

    it('has a list of stages as DOM child', () => {
      const children = pipeline[2];
      const stagesChild = children[0];

      expect(stagesChild[0]).to.equal(StageList);
      expect(stagesChild[1]).to.deep.equal([stageOne, stageTwo]);
    });

    it('has pipeline name as DOM child', () => {
      const children = pipeline[2];
      const nameChild = children[1];

      expect(nameChild[0]).to.equal(PipelineName);
      expect(nameChild[1]).to.equal('Dashy');
    });
  });
});
