// javascript/test/components/StageList.test.js
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2017. All rights reserved
// License::   MIT

var expect = require("chai").expect;

var Stage = require("../../src/components/Stage");
var StageList = require("../../src/components/StageList");

describe("StageList", function () {
    describe("#render", function () {
        var stageOne = {
            name: "Compile",
            status: "Passed"
        };
        var stageTwo = {
            name: "Test",
            status: "Building"
        };
        var stageList = new StageList().render([stageOne, stageTwo]);

        it("creates a DOM representation", function () {
            expect(stageList[0]).to.equal("div");
        });

        it("has CSS class", function () {
            expect(stageList[1]).to.deep.equal({
                class: "stage-container"
            });
        });

        it("has Stage children", function () {
            var children = stageList[2];
            expect(children).to.have.length(2);

            expect(children[0][0]).to.equal(Stage);
            expect(children[0][1]).to.equal(stageOne);

            expect(children[1][0]).to.equal(Stage);
            expect(children[1][1]).to.equal(stageTwo);
        });
    });
});
