// javascript/test/components/PipelineGroup.test.js
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2017. All rights reserved
// License::   MIT

const expect = require('chai').expect;

const Pipeline = require('../../src/components/Pipeline');
const PipelineGroup = require('../../src/components/PipelineGroup');

describe('PipelineGroup', () => {
  describe('#render', () => {
    const pipelineOneStages = [{
      name: 'DashyCompile',
      status: 'Passed',
    }, {
      name: 'DashyTest',
      status: 'Building',
    }];
    const pipelineTwoStages = [{
      name: 'FlashyCompile',
      status: 'Passed',
    }, {
      name: 'FlashyTest',
      status: 'Failing',
    }];

    const pipelineOne = {
      name: 'Dashy',
      stages: pipelineOneStages,
    };
    const pipelineTwo = {
      name: 'Dashy',
      stages: pipelineTwoStages,
    };

    const pipelineGroup = new PipelineGroup().render([pipelineOne, pipelineTwo]);

    it('has creates a DOM representation', () => {
      expect(pipelineGroup[0]).to.equal('div');
    });

    it('has CSS class', () => {
      expect(pipelineGroup[1]).to.deep.equal({
        class: 'pipeline-group pipeline-group-2',
      });
    });

    it('has pipelines as children', () => {
      const children = pipelineGroup[2];
      const firstChild = children[0];
      const secondChild = children[1];

      expect(firstChild[0]).to.equal(Pipeline);
      expect(firstChild[1]).to.equal(pipelineOne);

      expect(secondChild[0]).to.equal(Pipeline);
      expect(secondChild[1]).to.equal(pipelineTwo);
    });
  });
});
