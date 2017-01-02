// javascript/test/components/StageList.test.js
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2017. All rights reserved
// License::   MIT

var Stage = require("../../src/components/Stage");
var StageList = require("../../src/components/StageList");

describe("StageList", function() {
    describe("#render", function() {
        var stageOne = {
            name: "Compile",
            status: "Passed"
        };
        var stageTwo = {
            name: "Test",
            status: "Building"
        };
        var stageList = new StageList().render([stageOne, stageTwo]);

        it("creates a DOM representation", function() {
            expect(stageList[0]).toEqual("div");
        });

        it("has CSS class", function() {
            expect(stageList[1]).toEqual({
                class: "stage-container"
            });
        });

        it("has Stage children", function() {
            var children = stageList[2];
            expect(children.length).toEqual(2);

            expect(children[0][0]).toEqual(Stage);
            expect(children[0][1]).toEqual(stageOne);

            expect(children[1][0]).toEqual(Stage);
            expect(children[1][1]).toEqual(stageTwo);
        });
    });
});