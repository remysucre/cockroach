// Copyright 2018 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package tests

var hibernateBlocklists = blocklistsForVersion{
	{"v20.2", "hibernateBlockList20_2", hibernateBlockList20_2, "", nil},
	{"v21.1", "hibernateBlockList21_1", hibernateBlockList21_1, "hibernateIgnoreList21_1", hibernateIgnoreList21_1},
	{"v21.2", "hibernateBlockList21_2", hibernateBlockList21_2, "hibernateIgnoreList21_2", hibernateIgnoreList21_2},
	{"v22.1", "hibernateBlockList22_1", hibernateBlockList22_1, "hibernateIgnoreList22_1", hibernateIgnoreList22_1},
}

var hibernateSpatialBlocklists = blocklistsForVersion{
	{"v21.1", "hibernateSpatialBlockList21_1", hibernateSpatialBlockList21_1, "", nil},
	{"v21.2", "hibernateSpatialBlockList21_2", hibernateSpatialBlockList21_2, "", nil},
	{"v22.1", "hibernateSpatialBlockList22_1", hibernateSpatialBlockList22_1, "", nil},
}

// Please keep these lists alphabetized for easy diffing.
// After a failed run, an updated version of this blocklist should be available
// in the test log.
var hibernateSpatialBlockList22_1 = blocklist{}

var hibernateSpatialBlockList21_2 = blocklist{}

var hibernateSpatialBlockList21_1 = blocklist{}

var hibernateBlockList22_1 = hibernateBlockList21_2

var hibernateBlockList21_2 = blocklist{
	"org.hibernate.jpa.test.graphs.FetchGraphTest.testCollectionEntityGraph":                                                "unknown",
	"org.hibernate.jpa.test.packaging.PackagedEntityManagerTest.testConfiguration":                                          "unknown",
	"org.hibernate.jpa.test.packaging.PackagedEntityManagerTest.testDefaultPar":                                             "unknown",
	"org.hibernate.jpa.test.packaging.PackagedEntityManagerTest.testDefaultParForPersistence_1_0":                           "unknown",
	"org.hibernate.jpa.test.packaging.PackagedEntityManagerTest.testExcludeHbmPar":                                          "unknown",
	"org.hibernate.jpa.test.packaging.PackagedEntityManagerTest.testExplodedPar":                                            "unknown",
	"org.hibernate.jpa.test.packaging.PackagedEntityManagerTest.testExtendedEntityManager":                                  "unknown",
	"org.hibernate.jpa.test.packaging.PackagedEntityManagerTest.testExternalJar":                                            "unknown",
	"org.hibernate.jpa.test.packaging.PackagedEntityManagerTest.testListeners":                                              "unknown",
	"org.hibernate.jpa.test.packaging.PackagedEntityManagerTest.testListenersDefaultPar":                                    "unknown",
	"org.hibernate.jpa.test.packaging.PackagedEntityManagerTest.testORMFileOnMainAndExplicitJars":                           "unknown",
	"org.hibernate.jpa.test.packaging.PackagedEntityManagerTest.testRelativeJarReferences":                                  "unknown",
	"org.hibernate.jpa.test.packaging.PackagedEntityManagerTest.testSpacePar":                                               "unknown",
	"org.hibernate.jpa.test.packaging.ScannerTest.testCustomScanner":                                                        "unknown",
	"org.hibernate.test.annotations.onetomany.OrderByTest.testOrderByNoElement":                                             "unknown",
	"org.hibernate.test.annotations.onetomany.OrderByTest.testOrderByOneToManyWithJoinTable":                                "unknown",
	"org.hibernate.test.bytecode.enhancement.lazy.proxy.inlinedirtychecking.DirtyCheckPrivateUnMappedCollectionTest.testIt": "unknown",
	"org.hibernate.test.hql.BulkManipulationTest.testUpdateWithSubquery":                                                    "unknown",
}

var hibernateBlockList21_1 = blocklist{
	"org.hibernate.jpa.test.graphs.FetchGraphTest.testCollectionEntityGraph":                                                                                                                     "unknown",
	"org.hibernate.jpa.test.packaging.PackagedEntityManagerTest.testConfiguration":                                                                                                               "unknown",
	"org.hibernate.jpa.test.packaging.PackagedEntityManagerTest.testDefaultPar":                                                                                                                  "unknown",
	"org.hibernate.jpa.test.packaging.PackagedEntityManagerTest.testDefaultParForPersistence_1_0":                                                                                                "unknown",
	"org.hibernate.jpa.test.packaging.PackagedEntityManagerTest.testExcludeHbmPar":                                                                                                               "unknown",
	"org.hibernate.jpa.test.packaging.PackagedEntityManagerTest.testExplodedPar":                                                                                                                 "unknown",
	"org.hibernate.jpa.test.packaging.PackagedEntityManagerTest.testExtendedEntityManager":                                                                                                       "unknown",
	"org.hibernate.jpa.test.packaging.PackagedEntityManagerTest.testExternalJar":                                                                                                                 "unknown",
	"org.hibernate.jpa.test.packaging.PackagedEntityManagerTest.testListeners":                                                                                                                   "unknown",
	"org.hibernate.jpa.test.packaging.PackagedEntityManagerTest.testListenersDefaultPar":                                                                                                         "unknown",
	"org.hibernate.jpa.test.packaging.PackagedEntityManagerTest.testORMFileOnMainAndExplicitJars":                                                                                                "unknown",
	"org.hibernate.jpa.test.packaging.PackagedEntityManagerTest.testRelativeJarReferences":                                                                                                       "unknown",
	"org.hibernate.jpa.test.packaging.PackagedEntityManagerTest.testSpacePar":                                                                                                                    "unknown",
	"org.hibernate.jpa.test.packaging.ScannerTest.testCustomScanner":                                                                                                                             "unknown",
	"org.hibernate.jpa.test.query.QueryTest.testParameterCollectionSingletonParenthesesAndPositional":                                                                                            "unknown",
	"org.hibernate.jpa.test.transaction.CloseEntityManagerWithActiveTransactionTest.testRemoveThenCloseWithAnActiveTransaction":                                                                  "unknown",
	"org.hibernate.jpa.test.transaction.CloseEntityManagerWithActiveTransactionTest.testUpdateThenCloseWithAnActiveTransaction":                                                                  "unknown",
	"org.hibernate.test.annotations.access.AccessTest.testNonOverridenSubclass":                                                                                                                  "unknown",
	"org.hibernate.test.annotations.access.AccessTest.testOverridenSubclass":                                                                                                                     "unknown",
	"org.hibernate.test.annotations.access.jpa.AccessTest.testDefaultConfigurationModeIsInherited":                                                                                               "unknown",
	"org.hibernate.test.annotations.access.jpa.AccessTest.testEmbeddableUsesAccessStrategyOfContainingClass":                                                                                     "unknown",
	"org.hibernate.test.annotations.access.jpa.AccessTest.testNonOverridenSubclass":                                                                                                              "unknown",
	"org.hibernate.test.annotations.access.jpa.AccessTest.testOverridenSubclass":                                                                                                                 "unknown",
	"org.hibernate.test.annotations.entity.BasicHibernateAnnotationsTest.testSetSimpleValueTypeNameInSecondPass":                                                                                 "unknown",
	"org.hibernate.test.annotations.entity.Java5FeaturesTest.testEnums":                                                                                                                          "unknown",
	"org.hibernate.test.annotations.id.IdTest.testIdClass":                                                                                                                                       "unknown",
	"org.hibernate.test.annotations.id.sequences.IdTest.testIdClass":                                                                                                                             "unknown",
	"org.hibernate.test.annotations.onetomany.OrderByTest.testOrderByNoElement":                                                                                                                  "unknown",
	"org.hibernate.test.annotations.onetomany.OrderByTest.testOrderByOneToManyWithJoinTable":                                                                                                     "unknown",
	"org.hibernate.test.bytecode.enhancement.lazy.proxy.DeepInheritanceProxyTest.testLeafGetValueToInitialize":                                                                                   "unknown",
	"org.hibernate.test.bytecode.enhancement.lazy.proxy.DeepInheritanceProxyTest.testLeafSetValueToInitialize":                                                                                   "unknown",
	"org.hibernate.test.bytecode.enhancement.lazy.proxy.DeepInheritanceWithNonEntitiesProxyTest.testLeafGetValueInNonEntity":                                                                     "unknown",
	"org.hibernate.test.bytecode.enhancement.lazy.proxy.DeepInheritanceWithNonEntitiesProxyTest.testLeafGetValueToInitialize":                                                                    "unknown",
	"org.hibernate.test.bytecode.enhancement.lazy.proxy.DeepInheritanceWithNonEntitiesProxyTest.testLeafSetValueInNonEntity":                                                                     "unknown",
	"org.hibernate.test.bytecode.enhancement.lazy.proxy.DeepInheritanceWithNonEntitiesProxyTest.testLeafSetValueToInitialize":                                                                    "unknown",
	"org.hibernate.test.bytecode.enhancement.lazy.proxy.inlinedirtychecking.DirtyCheckPrivateUnMappedCollectionTest.testIt":                                                                      "unknown",
	"org.hibernate.test.bytecode.enhancement.lazy.proxy.inlinedirtychecking.ManyToOnePropertyAccessByFieldTest.testUpdate":                                                                       "unknown",
	"org.hibernate.test.cache.polymorphism.PolymorphicCacheTest.testPolymorphismAndCache":                                                                                                        "unknown",
	"org.hibernate.test.converter.literal.QueryLiteralTest.testIntegerWrapper":                                                                                                                   "unknown",
	"org.hibernate.test.discriminator.SimpleInheritanceTest.testAccessAsIncorrectSubclass":                                                                                                       "unknown",
	"org.hibernate.test.discriminator.SimpleInheritanceTest.testDiscriminatorSubclass":                                                                                                           "unknown",
	"org.hibernate.test.dynamicentity.interceptor.InterceptorDynamicEntityTest.testIt":                                                                                                           "unknown",
	"org.hibernate.test.dynamicentity.tuplizer.TuplizerDynamicEntityTest.testIt":                                                                                                                 "unknown",
	"org.hibernate.test.dynamicentity.tuplizer2.ImprovedTuplizerDynamicEntityTest.testIt":                                                                                                        "unknown",
	"org.hibernate.test.ecid.EmbeddedCompositeIdTest.testMerge":                                                                                                                                  "unknown",
	"org.hibernate.test.event.collection.association.bidirectional.onetomany.BidirectionalOneToManyBagSubclassCollectionEventTest.testDeleteParentAndChild":                                      "unknown",
	"org.hibernate.test.event.collection.association.bidirectional.onetomany.BidirectionalOneToManyBagSubclassCollectionEventTest.testDeleteParentButNotChild":                                   "unknown",
	"org.hibernate.test.event.collection.association.bidirectional.onetomany.BidirectionalOneToManyBagSubclassCollectionEventTest.testDeleteParentWithNoChildren":                                "unknown",
	"org.hibernate.test.event.collection.association.bidirectional.onetomany.BidirectionalOneToManyBagSubclassCollectionEventTest.testDeleteParentWithNullChildren":                              "unknown",
	"org.hibernate.test.event.collection.association.bidirectional.onetomany.BidirectionalOneToManyBagSubclassCollectionEventTest.testMoveAllChildrenToDifferentParent":                          "unknown",
	"org.hibernate.test.event.collection.association.bidirectional.onetomany.BidirectionalOneToManyBagSubclassCollectionEventTest.testMoveChildToDifferentParent":                                "unknown",
	"org.hibernate.test.event.collection.association.bidirectional.onetomany.BidirectionalOneToManyBagSubclassCollectionEventTest.testMoveCollectionToDifferentParent":                           "unknown",
	"org.hibernate.test.event.collection.association.bidirectional.onetomany.BidirectionalOneToManyBagSubclassCollectionEventTest.testMoveCollectionToDifferentParentFlushMoveToDifferentParent": "unknown",
	"org.hibernate.test.event.collection.association.bidirectional.onetomany.BidirectionalOneToManyBagSubclassCollectionEventTest.testSaveParentEmptyChildren":                                   "unknown",
	"org.hibernate.test.event.collection.association.bidirectional.onetomany.BidirectionalOneToManyBagSubclassCollectionEventTest.testUpdateParentNoneToOneChild":                                "unknown",
	"org.hibernate.test.event.collection.association.bidirectional.onetomany.BidirectionalOneToManyBagSubclassCollectionEventTest.testUpdateParentNoneToOneChildDiffCollection":                  "unknown",
	"org.hibernate.test.event.collection.association.bidirectional.onetomany.BidirectionalOneToManyBagSubclassCollectionEventTest.testUpdateParentNullToOneChild":                                "unknown",
	"org.hibernate.test.event.collection.association.bidirectional.onetomany.BidirectionalOneToManyBagSubclassCollectionEventTest.testUpdateParentNullToOneChildDiffCollection":                  "unknown",
	"org.hibernate.test.event.collection.association.bidirectional.onetomany.BidirectionalOneToManyBagSubclassCollectionEventTest.testUpdateParentOneChildDiffCollectionDiffChild":               "unknown",
	"org.hibernate.test.event.collection.association.bidirectional.onetomany.BidirectionalOneToManyBagSubclassCollectionEventTest.testUpdateParentOneChildDiffCollectionSameChild":               "unknown",
	"org.hibernate.test.event.collection.association.bidirectional.onetomany.BidirectionalOneToManyBagSubclassCollectionEventTest.testUpdateParentOneChildToNoneByClear":                         "unknown",
	"org.hibernate.test.event.collection.association.bidirectional.onetomany.BidirectionalOneToManyBagSubclassCollectionEventTest.testUpdateParentOneChildToNoneByRemove":                        "unknown",
	"org.hibernate.test.event.collection.association.bidirectional.onetomany.BidirectionalOneToManyBagSubclassCollectionEventTest.testUpdateParentOneToTwoChildren":                              "unknown",
	"org.hibernate.test.event.collection.association.bidirectional.onetomany.BidirectionalOneToManyBagSubclassCollectionEventTest.testUpdateParentOneToTwoSameChildren":                          "unknown",
	"org.hibernate.test.event.collection.association.bidirectional.onetomany.BidirectionalOneToManyBagSubclassCollectionEventTest.testUpdateParentTwoChildrenToOne":                              "unknown",
	"org.hibernate.test.hql.BulkManipulationTest.testUpdateWithSubquery":                                                                                                                         "unknown",
	"org.hibernate.test.inheritance.TransientOverrideAsPersistentMappedSuperclass.testFindBySubclass":                                                                                            "unknown",
	"org.hibernate.test.inheritance.TransientOverrideAsPersistentSingleTable.testFindBySubclass":                                                                                                 "unknown",
	"org.hibernate.test.inheritance.discriminator.MultiSingleTableLoadTest.testFetchJoinLoadMultipleHoldersWithDifferentSubtypes":                                                                "unknown",
	"org.hibernate.test.inheritance.discriminator.SingleTableRelationsTest.testLazyInitialization":                                                                                               "unknown",
	"org.hibernate.test.joinedsubclass.JoinedSubclassTest.testLockingJoinedSubclass":                                                                                                             "unknown",
	"org.hibernate.test.jpa.compliance.tck2_2.caching.InheritedCacheableTest.testOnlySubclassIsCached":                                                                                           "unknown",
	"org.hibernate.test.legacy.CustomSQLTest.testJoinedSubclass":                                                                                                                                 "unknown",
	"org.hibernate.test.legacy.FooBarTest.testVersioning":                                                                                                                                        "unknown",
	"org.hibernate.test.legacy.FumTest.testDeleteOwner":                                                                                                                                          "unknown",
	"org.hibernate.test.naturalid.inheritance.cache.InheritedNaturalIdNoCacheTest.testLoadExtendedByNormal":                                                                                      "unknown",
	"org.hibernate.test.where.annotations.EagerManyToOneFetchModeSelectWhereTest.testAssociatedWhereClause":                                                                                      "unknown",
}

var hibernateBlockList20_2 = blocklist{
	"org.hibernate.jpa.test.graphs.FetchGraphTest.testCollectionEntityGraph":                                                                                                                     "unknown",
	"org.hibernate.jpa.test.packaging.PackagedEntityManagerTest.testConfiguration":                                                                                                               "unknown",
	"org.hibernate.jpa.test.packaging.PackagedEntityManagerTest.testDefaultPar":                                                                                                                  "unknown",
	"org.hibernate.jpa.test.packaging.PackagedEntityManagerTest.testDefaultParForPersistence_1_0":                                                                                                "unknown",
	"org.hibernate.jpa.test.packaging.PackagedEntityManagerTest.testExcludeHbmPar":                                                                                                               "unknown",
	"org.hibernate.jpa.test.packaging.PackagedEntityManagerTest.testExplodedPar":                                                                                                                 "unknown",
	"org.hibernate.jpa.test.packaging.PackagedEntityManagerTest.testExtendedEntityManager":                                                                                                       "unknown",
	"org.hibernate.jpa.test.packaging.PackagedEntityManagerTest.testExternalJar":                                                                                                                 "unknown",
	"org.hibernate.jpa.test.packaging.PackagedEntityManagerTest.testListeners":                                                                                                                   "unknown",
	"org.hibernate.jpa.test.packaging.PackagedEntityManagerTest.testListenersDefaultPar":                                                                                                         "unknown",
	"org.hibernate.jpa.test.packaging.PackagedEntityManagerTest.testORMFileOnMainAndExplicitJars":                                                                                                "unknown",
	"org.hibernate.jpa.test.packaging.PackagedEntityManagerTest.testRelativeJarReferences":                                                                                                       "unknown",
	"org.hibernate.jpa.test.packaging.PackagedEntityManagerTest.testSpacePar":                                                                                                                    "unknown",
	"org.hibernate.jpa.test.packaging.ScannerTest.testCustomScanner":                                                                                                                             "unknown",
	"org.hibernate.jpa.test.query.QueryTest.testParameterCollectionSingletonParenthesesAndPositional":                                                                                            "unknown",
	"org.hibernate.jpa.test.transaction.CloseEntityManagerWithActiveTransactionTest.testRemoveThenCloseWithAnActiveTransaction":                                                                  "unknown",
	"org.hibernate.jpa.test.transaction.CloseEntityManagerWithActiveTransactionTest.testUpdateThenCloseWithAnActiveTransaction":                                                                  "unknown",
	"org.hibernate.test.annotations.access.AccessTest.testNonOverridenSubclass":                                                                                                                  "unknown",
	"org.hibernate.test.annotations.access.AccessTest.testOverridenSubclass":                                                                                                                     "unknown",
	"org.hibernate.test.annotations.access.jpa.AccessTest.testDefaultConfigurationModeIsInherited":                                                                                               "unknown",
	"org.hibernate.test.annotations.access.jpa.AccessTest.testEmbeddableUsesAccessStrategyOfContainingClass":                                                                                     "unknown",
	"org.hibernate.test.annotations.access.jpa.AccessTest.testNonOverridenSubclass":                                                                                                              "unknown",
	"org.hibernate.test.annotations.access.jpa.AccessTest.testOverridenSubclass":                                                                                                                 "unknown",
	"org.hibernate.test.annotations.entity.BasicHibernateAnnotationsTest.testSetSimpleValueTypeNameInSecondPass":                                                                                 "unknown",
	"org.hibernate.test.annotations.entity.Java5FeaturesTest.testEnums":                                                                                                                          "unknown",
	"org.hibernate.test.annotations.id.IdTest.testIdClass":                                                                                                                                       "unknown",
	"org.hibernate.test.annotations.id.sequences.IdTest.testIdClass":                                                                                                                             "unknown",
	"org.hibernate.test.annotations.onetomany.OrderByTest.testOrderByNoElement":                                                                                                                  "unknown",
	"org.hibernate.test.annotations.onetomany.OrderByTest.testOrderByOneToManyWithJoinTable":                                                                                                     "unknown",
	"org.hibernate.test.bytecode.enhancement.lazy.proxy.DeepInheritanceProxyTest.testLeafGetValueToInitialize":                                                                                   "unknown",
	"org.hibernate.test.bytecode.enhancement.lazy.proxy.DeepInheritanceProxyTest.testLeafSetValueToInitialize":                                                                                   "unknown",
	"org.hibernate.test.bytecode.enhancement.lazy.proxy.DeepInheritanceWithNonEntitiesProxyTest.testLeafGetValueInNonEntity":                                                                     "unknown",
	"org.hibernate.test.bytecode.enhancement.lazy.proxy.DeepInheritanceWithNonEntitiesProxyTest.testLeafGetValueToInitialize":                                                                    "unknown",
	"org.hibernate.test.bytecode.enhancement.lazy.proxy.DeepInheritanceWithNonEntitiesProxyTest.testLeafSetValueInNonEntity":                                                                     "unknown",
	"org.hibernate.test.bytecode.enhancement.lazy.proxy.DeepInheritanceWithNonEntitiesProxyTest.testLeafSetValueToInitialize":                                                                    "unknown",
	"org.hibernate.test.bytecode.enhancement.lazy.proxy.inlinedirtychecking.DirtyCheckPrivateUnMappedCollectionTest.testIt":                                                                      "unknown",
	"org.hibernate.test.bytecode.enhancement.lazy.proxy.inlinedirtychecking.ManyToOnePropertyAccessByFieldTest.testUpdate":                                                                       "unknown",
	"org.hibernate.test.cache.polymorphism.PolymorphicCacheTest.testPolymorphismAndCache":                                                                                                        "unknown",
	"org.hibernate.test.converter.literal.QueryLiteralTest.testIntegerWrapper":                                                                                                                   "unknown",
	"org.hibernate.test.discriminator.SimpleInheritanceTest.testAccessAsIncorrectSubclass":                                                                                                       "unknown",
	"org.hibernate.test.discriminator.SimpleInheritanceTest.testDiscriminatorSubclass":                                                                                                           "unknown",
	"org.hibernate.test.dynamicentity.interceptor.InterceptorDynamicEntityTest.testIt":                                                                                                           "unknown",
	"org.hibernate.test.dynamicentity.tuplizer.TuplizerDynamicEntityTest.testIt":                                                                                                                 "unknown",
	"org.hibernate.test.dynamicentity.tuplizer2.ImprovedTuplizerDynamicEntityTest.testIt":                                                                                                        "unknown",
	"org.hibernate.test.ecid.EmbeddedCompositeIdTest.testMerge":                                                                                                                                  "unknown",
	"org.hibernate.test.event.collection.association.bidirectional.onetomany.BidirectionalOneToManyBagSubclassCollectionEventTest.testDeleteParentAndChild":                                      "unknown",
	"org.hibernate.test.event.collection.association.bidirectional.onetomany.BidirectionalOneToManyBagSubclassCollectionEventTest.testDeleteParentButNotChild":                                   "unknown",
	"org.hibernate.test.event.collection.association.bidirectional.onetomany.BidirectionalOneToManyBagSubclassCollectionEventTest.testDeleteParentWithNoChildren":                                "unknown",
	"org.hibernate.test.event.collection.association.bidirectional.onetomany.BidirectionalOneToManyBagSubclassCollectionEventTest.testDeleteParentWithNullChildren":                              "unknown",
	"org.hibernate.test.event.collection.association.bidirectional.onetomany.BidirectionalOneToManyBagSubclassCollectionEventTest.testMoveAllChildrenToDifferentParent":                          "unknown",
	"org.hibernate.test.event.collection.association.bidirectional.onetomany.BidirectionalOneToManyBagSubclassCollectionEventTest.testMoveChildToDifferentParent":                                "unknown",
	"org.hibernate.test.event.collection.association.bidirectional.onetomany.BidirectionalOneToManyBagSubclassCollectionEventTest.testMoveCollectionToDifferentParent":                           "unknown",
	"org.hibernate.test.event.collection.association.bidirectional.onetomany.BidirectionalOneToManyBagSubclassCollectionEventTest.testMoveCollectionToDifferentParentFlushMoveToDifferentParent": "unknown",
	"org.hibernate.test.event.collection.association.bidirectional.onetomany.BidirectionalOneToManyBagSubclassCollectionEventTest.testSaveParentEmptyChildren":                                   "unknown",
	"org.hibernate.test.event.collection.association.bidirectional.onetomany.BidirectionalOneToManyBagSubclassCollectionEventTest.testUpdateParentNoneToOneChild":                                "unknown",
	"org.hibernate.test.event.collection.association.bidirectional.onetomany.BidirectionalOneToManyBagSubclassCollectionEventTest.testUpdateParentNoneToOneChildDiffCollection":                  "unknown",
	"org.hibernate.test.event.collection.association.bidirectional.onetomany.BidirectionalOneToManyBagSubclassCollectionEventTest.testUpdateParentNullToOneChild":                                "unknown",
	"org.hibernate.test.event.collection.association.bidirectional.onetomany.BidirectionalOneToManyBagSubclassCollectionEventTest.testUpdateParentNullToOneChildDiffCollection":                  "unknown",
	"org.hibernate.test.event.collection.association.bidirectional.onetomany.BidirectionalOneToManyBagSubclassCollectionEventTest.testUpdateParentOneChildDiffCollectionDiffChild":               "unknown",
	"org.hibernate.test.event.collection.association.bidirectional.onetomany.BidirectionalOneToManyBagSubclassCollectionEventTest.testUpdateParentOneChildDiffCollectionSameChild":               "unknown",
	"org.hibernate.test.event.collection.association.bidirectional.onetomany.BidirectionalOneToManyBagSubclassCollectionEventTest.testUpdateParentOneChildToNoneByClear":                         "unknown",
	"org.hibernate.test.event.collection.association.bidirectional.onetomany.BidirectionalOneToManyBagSubclassCollectionEventTest.testUpdateParentOneChildToNoneByRemove":                        "unknown",
	"org.hibernate.test.event.collection.association.bidirectional.onetomany.BidirectionalOneToManyBagSubclassCollectionEventTest.testUpdateParentOneToTwoChildren":                              "unknown",
	"org.hibernate.test.event.collection.association.bidirectional.onetomany.BidirectionalOneToManyBagSubclassCollectionEventTest.testUpdateParentOneToTwoSameChildren":                          "unknown",
	"org.hibernate.test.event.collection.association.bidirectional.onetomany.BidirectionalOneToManyBagSubclassCollectionEventTest.testUpdateParentTwoChildrenToOne":                              "unknown",
	"org.hibernate.test.hql.BulkManipulationTest.testUpdateWithSubquery":                                                                                                                         "unknown",
	"org.hibernate.test.inheritance.TransientOverrideAsPersistentMappedSuperclass.testFindBySubclass":                                                                                            "unknown",
	"org.hibernate.test.inheritance.TransientOverrideAsPersistentSingleTable.testFindBySubclass":                                                                                                 "unknown",
	"org.hibernate.test.inheritance.discriminator.MultiSingleTableLoadTest.testFetchJoinLoadMultipleHoldersWithDifferentSubtypes":                                                                "unknown",
	"org.hibernate.test.inheritance.discriminator.SingleTableRelationsTest.testLazyInitialization":                                                                                               "unknown",
	"org.hibernate.test.joinedsubclass.JoinedSubclassTest.testLockingJoinedSubclass":                                                                                                             "unknown",
	"org.hibernate.test.jpa.compliance.tck2_2.caching.InheritedCacheableTest.testOnlySubclassIsCached":                                                                                           "unknown",
	"org.hibernate.test.legacy.CustomSQLTest.testJoinedSubclass":                                                                                                                                 "unknown",
	"org.hibernate.test.legacy.FooBarTest.testVersioning":                                                                                                                                        "unknown",
	"org.hibernate.test.legacy.FumTest.testDeleteOwner":                                                                                                                                          "unknown",
	"org.hibernate.test.naturalid.inheritance.cache.InheritedNaturalIdNoCacheTest.testLoadExtendedByNormal":                                                                                      "unknown",
	"org.hibernate.test.where.annotations.EagerManyToOneFetchModeSelectWhereTest.testAssociatedWhereClause":                                                                                      "unknown",
}

var hibernateIgnoreList22_1 = hibernateIgnoreList21_2

var hibernateIgnoreList21_2 = hibernateIgnoreList21_1

var hibernateIgnoreList21_1 = blocklist{
	"org.hibernate.userguide.pc.WhereTest.testLifecycle": "unknown",
}
