// javascript/test/components/PipelineName.test.js
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2017. All rights reserved
// License::   MIT

var expect = require("chai").expect;

var PipelineName = require("../../src/components/PipelineName");

describe("PipelineName", function () {
    describe("#render", function () {
        var stage = new PipelineName().render("Dashy");

        it("creates a DOM representation", function () {
            expect(stage[0]).to.equal("div");
        });

        it("has CSS class", function () {
            expect(stage[1]).to.deep.equal({
                class: "pipeline-name"
            });
        });

        it("has contents based on its name", function () {
            expect(stage[2]).to.equal("Dashy");
        });
    });
});
