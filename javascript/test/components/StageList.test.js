var Stage = require("../../src/components/Stage");
var StageList = require("../../src/components/StageList");

describe("StageList", function() {
    describe("#render", function() {
        var stageList = new StageList().render([{
            "name": "Compile",
            "status": "Passed"
        }, {
            "name": "Test",
            "status": "Building"
        }]);

        it("creates a DOM representation", function() {
            expect(stageList[0]).toEqual("div");
        });

        it("has CSS class", function() {
            expect(stageList[1]).toEqual({
                "class": "stage-container"
            });
        });

        it("has Stage children", function() {
            var children = stageList[2];
            expect(children.length).toEqual(2);

            var firstChild = children[0];
            expect(firstChild[0]).toEqual(Stage);
            expect(firstChild[1]).toEqual({
                "name": "Compile",
                "status": "Passed"
            });

            var secondChild = children[1];
            expect(secondChild[0]).toEqual(Stage);
            expect(secondChild[1]).toEqual({
                "name": "Test",
                "status": "Building"
            });
        });
    });
});
